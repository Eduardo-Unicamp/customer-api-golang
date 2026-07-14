package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"first-api/internal/auth"
	"first-api/internal/model" // Importe seu repositório de clientes
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CustomerAuthRepository interface {
	GetCustomers(context.Context) ([]model.Customer, error)
	CreateCustomer(context.Context, *model.Customer) error
	GetCustomerByField(context.Context, string, string) (*model.Customer, error)
}

type AuthUseCase struct {
	CustomerRepo CustomerAuthRepository
	jwtConfig    *auth.JWTConfig
}

func NewAuthUseCase(cr CustomerAuthRepository, config *auth.JWTConfig) *AuthUseCase {
	return &AuthUseCase{
		CustomerRepo: cr,
		jwtConfig:    config,
	}
}

func (au *AuthUseCase) Register(ctx context.Context, r *http.Request) (*model.TokenResponseDTO, error) {
	var request model.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("Error while parsing the json:%w", err)
	}

	_, err := au.CustomerRepo.GetCustomerByField(ctx, "email", request.Email)
	if err == nil {
		return nil, model.ErrEmailTaken
	} //achou, email repetido
	if !errors.Is(err, model.CustomerNotFound) {
		return nil, err
	} //se deu algum outro erro

	customer, err := model.NewCustomer(request.Name, request.Email, request.Phone, request.Password)
	if err != nil {
		return nil, err
	}

	if err := au.CustomerRepo.CreateCustomer(ctx, customer); err != nil {
		return nil, err
	}

	return au.GenerateTokenResponse(customer.ID)

}

func (au *AuthUseCase) Login(ctx context.Context, r *http.Request) (*model.TokenResponseDTO, error) {
	var request model.LoginDTO
	json.NewDecoder(r.Body).Decode(&request)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, model.ErrReadingJSON
	}

	customer, err := au.CustomerRepo.GetCustomerByField(ctx, "email", request.Email)
	if err != nil {
		return nil, model.ErrInvalidPassword //sim, é email, mas tem aquela regrinha de nao dizer se é email ou senha que errou
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password))
	if err != nil {
		return nil, model.ErrInvalidPassword
	}

	return au.GenerateTokenResponse(customer.ID)
}

func (au *AuthUseCase) GenerateTokenResponse(customerID uuid.UUID) (*model.TokenResponseDTO, error) {
	tokenStr, err := auth.GenerateToken(customerID, au.jwtConfig)
	if err != nil {
		return nil, err
	}

	return &model.TokenResponseDTO{
		AccessToken: tokenStr,
		ExpiresIn:   au.jwtConfig.ExpirationMinutes * 60,
		CustomerID:  customerID,
	}, nil
}
