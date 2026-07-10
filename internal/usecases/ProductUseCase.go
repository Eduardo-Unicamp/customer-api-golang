package usecases

import (
	"context"
	"encoding/json"
	"first-api/internal/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProductRepository interface {
	GetProducts(context.Context) (*[]model.Product, error)
	GetProductByID(context.Context, string) (*model.Product, error)
	CreateProduct(context.Context, *model.Product) error
	UpdateProduct(context.Context, string, *model.Product) error
	DeleteProduct(context.Context, string) error
}

type ProductUseCase struct {
	repository ProductRepository
}

func NewProductUseCase(repository ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repository: repository,
	}
}

func (pu *ProductUseCase) GetProducts(ctx context.Context) (*[]model.Product, error) {
	products, err := pu.repository.GetProducts(ctx)
	if err != nil {
		return &[]model.Product{}, err
	}

	return products, err
}

func (pu *ProductUseCase) GetProductByID(ctx context.Context, r *http.Request) (*model.Product, error) {
	productID := chi.URLParam(r, "product_id")
	product, err := pu.repository.GetProductByID(ctx, productID)

	return product, err

}

func (pu *ProductUseCase) CreateProduct(ctx context.Context, r *http.Request) (*model.Product, error) {
	var request model.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	product, err := model.NewProduct(request.Name, request.Price, request.Stock)
	if err != nil {
		return nil, err
	}

	if err := pu.repository.CreateProduct(ctx, product); err != nil {
		return nil, err
	}

	return product, nil

}

func (pu *ProductUseCase) UpdateProduct(ctx context.Context, r *http.Request) (*model.Product, error) {
	productId := chi.URLParam(r, "productId")
	var request model.UpdateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return &model.Product{}, err
	}

	product, err := model.NewProduct(request.Name, request.Price, request.Stock)
	if err != nil {
		return product, err
	}

	err = pu.repository.UpdateProduct(ctx, productId, product)

	return product, err

}

func (pu *ProductUseCase) DeleteProduct(ctx context.Context, r *http.Request) error {
	productId := chi.URLParam(r, "productId")
	err := pu.repository.DeleteProduct(ctx, productId)
	if err != nil {
		return err
	}
	return nil

}
