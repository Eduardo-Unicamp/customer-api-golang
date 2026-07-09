package model

import (
	"errors"
)

var ErrNameRequired = errors.New("O nome não pode ser vazio")

var ErrEmailRequired = errors.New("O email não pode ser vazio")

var ErrPhoneRequired = errors.New("O telefone não pode ser vazio")

var CustomerNotFound = errors.New("Cliente não encontrado")

var ErrEmptyString = errors.New("String vazia")

var ErrInvalidPrice = errors.New("O preço informado não é válido")

var ErrInvalidStockQuantity = errors.New("Quantidade de estoque informada não é válida")

var ErrEmptyOrder = errors.New("O pedido precisa conter ao menos um item")

var ErrInvalidOrderStatus = errors.New("Status de pedido inválido.")
