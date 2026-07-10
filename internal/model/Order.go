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
	BoughtPrice  decimal.Decimal `json:"price"`
	UnitsOrdered int             `json:"amount"`
}

func NewOrder(customerID uuid.UUID, items *[]OrderItem) (*Order, error) {
	if len(*items) < 1 {
		return &Order{}, ErrEmptyOrder
	}

	var err error
	order := Order{}
	//gera id
	order.ID, err = uuid.NewV7()
	if err != nil {
		return &order, err
	}

	order.Items = *items
	order.Status = PENDING //"o produto deve nascer como PENDING" segundo as intruções do projeto no github

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
	orderItem.BoughtPrice = price

	err = ValidateStockQuantity(amount)
	if err != nil {
		return &OrderItem{}, err
	}
	orderItem.UnitsOrdered = amount

	return &orderItem, nil

}

func (o *Order) pay() {
	o.Status = PAID
}

func (o *Order) cancelOrder() {
	o.Status = CANCELED

}
