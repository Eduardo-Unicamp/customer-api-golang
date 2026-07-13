package handler

import (
	"context"

	"encoding/json"

	"first-api/internal/model"

	"net/http"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, r *http.Request) (*model.Order, error)
	GetOrders(ctx context.Context, r *http.Request) (*[]model.Order, error)
	GetOrderByID(ctx context.Context, r *http.Request) (*model.Order, error)
	PayOrder(ctx context.Context, r *http.Request) error
	CancelOrder(ctx context.Context, r *http.Request) error
}

type OrderHandler struct {
	UseCase OrderUseCase
}

func NewOrderHandler(orderUseCase OrderUseCase) *OrderHandler {
	return &OrderHandler{UseCase: orderUseCase}
}

func (oh *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	order, err := oh.UseCase.CreateOrder(ctx, r)
	if err != nil {
		WriteOrderError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(*order)

}

func (oh *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	orders, err := oh.UseCase.GetOrders(ctx, r)
	if err != nil {
		WriteOrderError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*orders)

}

func (oh *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	order, err := oh.UseCase.GetOrderByID(ctx, r)
	if err != nil {
		WriteOrderError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*order)

}

func (oh *OrderHandler) PayOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := oh.UseCase.PayOrder(ctx, r); err != nil {
		WriteOrderError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func (oh *OrderHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := oh.UseCase.CancelOrder(ctx, r); err != nil {
		WriteOrderError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
