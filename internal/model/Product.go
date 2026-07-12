package model

import (
	"strings"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID    uuid.UUID       `json:"id"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
	Stock int             `json:"stock"`
}

func NewProduct(name string, price decimal.Decimal, stock int) (*Product, error) {
	var err error
	//id
	product := Product{}
	if product.ID, err = uuid.NewV7(); err != nil {
		return &Product{}, err
	}
	//name
	if product.Name, err = ValidateString(name); err != nil {
		return &Product{}, err
	}
	//price
	if err := ValidatePrice(price); err != nil {
		return &Product{}, err
	}
	product.Price = price

	//stock
	if err := ValidateStockQuantity(stock); err != nil {
		return &Product{}, err
	}
	product.Stock = stock

	return &product, nil

}

type CreateProductRequest struct {
	ID    uuid.UUID       `json:"id"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
	Stock int             `json:"stock"`
}

type UpdateProductRequest struct {
	ID    uuid.UUID       `json:"id"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
	Stock int             `json:"stock"`
}

// to get product info for orders actions
type ProductInfo struct {
	CurrentPrices map[string]decimal.Decimal
	CurrentStock  map[string]int
}

func NewProductInfo() *ProductInfo {
	idPrice := make(map[string]decimal.Decimal)
	idStock := make(map[string]int)
	return &ProductInfo{CurrentPrices: idPrice, CurrentStock: idStock}
}

func ValidateString(target string) (string, error) {
	target = strings.TrimSpace(target)
	if target == "" {
		return target, ErrEmptyString
	}
	return target, nil
}

func ValidatePrice(price decimal.Decimal) error {
	//price must be positive
	if price.Compare(decimal.Zero) < 0 {
		return ErrInvalidPrice
	}

	//everything passed
	return nil
}
func ValidateStockQuantity(stockQuantity int) error {
	if stockQuantity <= 0 {
		return ErrInvalidStockQuantity
	}

	return nil
}
