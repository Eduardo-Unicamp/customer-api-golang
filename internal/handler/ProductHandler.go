package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"first-api/internal/model"
)

type ProductUseCase interface {
	GetProducts(context.Context) (*[]model.Product, error)
	GetProductByID(context.Context, *http.Request) (*model.Product, error)
	CreateProduct(context.Context, *http.Request) (*model.Product, error)
	UpdateProduct(context.Context, *http.Request) (*model.Product, error)
	DeleteProduct(context.Context, *http.Request) error
}

type ProductHandler struct {
	UseCase ProductUseCase
}

func NewProductHandler(useCase ProductUseCase) *ProductHandler {
	return &ProductHandler{UseCase: useCase}
}

func (p *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	products, err := p.UseCase.GetProducts(ctx)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (p *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	product, err := p.UseCase.GetProductByID(ctx, r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	product, err := p.UseCase.CreateProduct(ctx, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*product)

}

func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	product, err := p.UseCase.UpdateProduct(ctx, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*product)
}

func (p *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := p.UseCase.DeleteProduct(ctx, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
