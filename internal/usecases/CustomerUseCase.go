package usecases

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"first-api/internal/model"

	"net/http"

	"github.com/go-chi/chi/v5"
)

type CustomerRepository interface {
	GetCustomers(context.Context) ([]model.Customer, error)
	CreateCustomer(context.Context, *model.Customer) error
	UpdateCustomer(context.Context, string, *model.Customer) error
	DeleteCustomer(context.Context, string) error
	GetCustomerByField(context.Context, string, string) (*model.Customer, error)
}

type CustomerUseCase struct {
	repository CustomerRepository
}

func NewCustomerUseCase(repository CustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{
		repository: repository,
	}
}

func (pu *CustomerUseCase) GetCustomers(ctx context.Context) ([]model.Customer, error) {
	customers, err := pu.repository.GetCustomers(ctx)
	if err != nil {
		return []model.Customer{}, err
	}

	return customers, err
}

func (pu *CustomerUseCase) CreateCustomer(ctx context.Context, r *http.Request) (*model.Customer, error) {
	var request model.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	//check for duplicated email
	_, err := pu.repository.GetCustomerByField(ctx, "email", request.Email)
	if err == nil { //aqui se deu certo já tem esse email
		return nil, model.ErrEmailTaken
	}
	if !errors.Is(err, sql.ErrNoRows) { //se erro não é nil(passou do outro if) e nao é NoRows é pq deu algum outro erro aí retorna
		return nil, err
	}

	customer, err := model.NewCustomer(request.Name, request.Email, request.Phone, request.Password)
	if err != nil {
		return nil, err
	}

	if err := pu.repository.CreateCustomer(ctx, customer); err != nil {
		return nil, err
	}

	return customer, nil

}

func (cu *CustomerUseCase) UpdateCustomer(ctx context.Context, r *http.Request) (*model.Customer, error) {
	customerId := chi.URLParam(r, "customerId")
	var request model.UpdateCustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return &model.Customer{}, err
	}

	customer, err := model.NewCustomer(request.Name, request.Email, request.Phone, request.Password)
	if err != nil {
		return customer, err
	}

	err = cu.repository.UpdateCustomer(ctx, customerId, customer)

	return customer, err

}

func (pu *CustomerUseCase) DeleteCustomer(ctx context.Context, r *http.Request) error {
	customerId := chi.URLParam(r, "customerId")
	err := pu.repository.DeleteCustomer(ctx, customerId)
	if err != nil {
		return err
	}
	return nil

}
