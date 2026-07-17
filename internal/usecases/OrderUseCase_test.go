package usecases

import (
	"context"
	"database/sql"
	"first-api/internal/mocks"
	"first-api/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetOrderByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOrderRepository := mocks.NewMockOrderRepository(ctrl)
	mockProductRepository := mocks.NewMockProductRepositoryForOrder(ctrl)
	orderUseCase := NewOrderUseCase(mockOrderRepository, mockProductRepository)

	validUUID := uuid.MustParse("019f5936-935b-7e2a-985e-468106c73243")
	invalidUUID := "invalidUUID"

	testOrder := model.Order{
		ID:         validUUID,
		CustomerID: validUUID,
		Items: []model.OrderItem{
			{
				ID:           validUUID,
				ProductID:    validUUID,
				SellingPrice: decimal.NewFromInt(1),
				UnitsOrdered: 1,
			},
		},
		Status: model.PENDING,
	}

	//Arrange
	testCases := []struct {
		name           string
		id             string
		setupMocks     func()
		expectedResult *model.Order
		expectedError  error
	}{
		{
			name: "SUCCESS Id válido, retorna order",
			id:   validUUID.String(),
			setupMocks: func() {
				mockOrderRepository.EXPECT().GetOrderByID(gomock.Any(), "019f5936-935b-7e2a-985e-468106c73243").
					Return(&testOrder, nil).Times(1)
			},
			expectedResult: &testOrder,
			expectedError:  nil,
		},
		{
			name:           "Id em formato inválido, retorna erro",
			id:             invalidUUID,
			setupMocks:     func() {},
			expectedResult: nil,
			expectedError:  model.InvalidIDFormat,
		},
		{
			name: "Id em formato válido, mas order inexistente",
			id:   "019f5936-935b-7e2a-985e-468106c73243", //not in the database
			setupMocks: func() {
				mockOrderRepository.EXPECT().GetOrderByID(gomock.Any(), "019f5936-935b-7e2a-985e-468106c73243").Return(nil, sql.ErrNoRows).Times(1)
			},
			expectedResult: nil,
			expectedError:  model.ErrOrderNotFound,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			//Arrange
			tt.setupMocks()

			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("order_id", tt.id)
			r := httptest.NewRequest(http.MethodGet, "/order/{order_id}", nil)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

			//Act
			order, err := orderUseCase.GetOrderByID(context.Background(), r)

			//Assert
			assert.Equal(t, tt.expectedResult, order, "")

			if tt.expectedError != nil {
				assert.Error(t, err) //se esperava erro assegura que ele ocorreu, se nao assegura q ele nao ocorreu
			} else {
				assert.NoError(t, err)
			}

		})
	}

}
