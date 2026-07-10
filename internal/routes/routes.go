package routes

import (
	"first-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func CustomerRoutes(r *chi.Mux, handler *handler.CustomerHandler) {
	r.Get("/customer", handler.GetCustomers)
	r.Post("/customer", handler.CreateCustomer)
	r.Put("/customer/{customerId}", handler.UpdateCustomer)
	r.Delete("/customer/{customerId}", handler.DeleteCustomer)

}

func ProductRoutes(r *chi.Mux, handler *handler.ProductHandler) {
	r.Get("/product", handler.GetProducts)
	r.Post("/product", handler.CreateProduct)
	r.Put("/product/{customerId}", handler.UpdateProduct)
	r.Delete("/product/{customerId}", handler.DeleteProduct)
}
