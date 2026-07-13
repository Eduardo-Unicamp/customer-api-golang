package routes

import (
	"first-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func CustomerRoutes(r *chi.Mux, handler *handler.CustomerHandler) {
	r.Get("/customer", handler.GetCustomers)
	r.Get("/customer/{customer_id}", handler.GetCustomerByID)
	r.Post("/customer", handler.CreateCustomer)
	r.Put("/customer/{customer_id}", handler.UpdateCustomer)
	r.Delete("/customer/{customer_id}", handler.DeleteCustomer)

}

func ProductRoutes(r *chi.Mux, handler *handler.ProductHandler) {
	r.Get("/product", handler.GetProducts)
	r.Get("/product/{product_id}", handler.GetProductByID)
	r.Post("/product", handler.CreateProduct)
	r.Put("/product/{product_id}", handler.UpdateProduct)
	r.Delete("/product/{product_id}", handler.DeleteProduct)
}

func OrderRoutes(r *chi.Mux, handler *handler.OrderHandler) {
	r.Get("/order/all/{customer_id}", handler.GetOrders)
	r.Get("/order/{order_id}", handler.GetOrderByID)
	r.Post("/order", handler.CreateOrder)
	r.Post("/order/cancel/{order_id}", handler.CancelOrder)
	r.Post("/order/pay/{order_id}", handler.PayOrder)
}
