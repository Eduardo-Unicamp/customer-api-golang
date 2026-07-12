package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Order struct {
	ID         uuid.UUID   `json:"id"`
	CustomerID uuid.UUID   `json:"customer_id"`
	Items      []OrderItem `json:"items"`
	Status     OrderStatus `json:"status"`
}

type OrderItem struct {
	ID           uuid.UUID       `json:"id"`
	ProductID    uuid.UUID       `json:"product_id"`
	SellingPrice decimal.Decimal `json:"price"`
	UnitsOrdered int             `json:"amount"`
}

func NewOrder(customerID uuid.UUID, itemsDTOs []NewOrderItemDTO, productPrices map[string]decimal.Decimal) (*Order, error) {
	if len(itemsDTOs) < 1 {
		return &Order{}, ErrEmptyOrder
	}

	var err error
	order := Order{}
	//gera id
	order.ID, err = uuid.NewV7()
	if err != nil {
		return &order, err
	}
	order.CustomerID = customerID
	order.Status = PENDING //"o produto deve nascer como PENDING" segundo as intruções do projeto no github

	for _, itemDTO := range itemsDTOs {
		currentPrice := productPrices[itemDTO.ProductID.String()]
		newItem, err := NewOrderItem(itemDTO.ProductID, currentPrice, itemDTO.UnitsOrdered)
		if err != nil {
			return &Order{}, err
		}
		order.Items = append(order.Items, *newItem)
	}

	return &order, nil
}

func NewOrderItem(productID uuid.UUID, price decimal.Decimal, amount int) (*OrderItem, error) {
	var err error
	orderItem := OrderItem{}

	orderItem.ID, err = uuid.NewV7()
	if err != nil {
		return &orderItem, err
	}

	if uuid.Validate(productID.String()) != nil {
		return &OrderItem{}, err
	}
	orderItem.ProductID = productID

	err = ValidatePrice(price)
	if err != nil {
		return &OrderItem{}, err
	}
	orderItem.SellingPrice = price

	err = ValidateStockQuantity(amount)
	if err != nil {
		return &OrderItem{}, err
	}
	orderItem.UnitsOrdered = amount

	return &orderItem, nil

}

func (o *Order) Pay() {
	o.Status = PAID
}

func (o *Order) Cancel() {
	o.Status = CANCELED

}

type NewOrderDTO struct {
	CustomerID uuid.UUID         `json:"customer_id"`
	Items      []NewOrderItemDTO `json:"items"`
}

type NewOrderItemDTO struct {
	ProductID    uuid.UUID `json:"product_id"`
	UnitsOrdered int       `json:"amount"`
}
