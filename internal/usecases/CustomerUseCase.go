package usecases

import (
	"context"
	"database/sql"
	"errors"
	"first-api/internal/middleware"
	"first-api/internal/model"
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

func (pu *CustomerUseCase) GetCustomerByID(ctx context.Context) (*model.Customer, error) {
	customerID := middleware.GetUserIDFromToken(ctx)
	customer, err := pu.repository.GetCustomerByField(ctx, "id", customerID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (pu *CustomerUseCase) CreateCustomer(ctx context.Context, request model.CreateCustomerRequest) (*model.Customer, error) {

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

func (cu *CustomerUseCase) UpdateCustomer(ctx context.Context, request model.UpdateCustomerRequest) (*model.Customer, error) {
	customerID := middleware.GetUserIDFromToken(ctx)

	customer, err := model.NewCustomer(request.Name, request.Email, request.Phone, request.Password)
	if err != nil {
		return customer, err
	}

	err = cu.repository.UpdateCustomer(ctx, customerID, customer)

	return customer, err

}

func (pu *CustomerUseCase) DeleteCustomer(ctx context.Context) error {
	customerID := middleware.GetUserIDFromToken(ctx)
	err := pu.repository.DeleteCustomer(ctx, customerID)
	if err != nil {
		return err
	}
	return nil

}
