package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"first-api/db"
	"first-api/internal/handler"
	"first-api/internal/repository"
	"first-api/internal/routes"
	"first-api/internal/usecases"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar as variáveis de ambiente: %s", err)
	}

	dbPool, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer dbPool.Close()

	CustomerRepository := repository.NewCustomerRepository(dbPool)
	CustomerUseCase := usecases.NewCustomerUseCase(CustomerRepository)
	CustomerHandler := handler.NewCustomerHandler(CustomerUseCase)

	ProductRepository := repository.NewProductRepository(dbPool)
	ProductUseCase := usecases.NewProductUseCase(ProductRepository)
	ProductHandler := handler.NewProductHandler(ProductUseCase)

	OrderRepository := repository.NewOrderRepository(dbPool)
	OrderUseCase := usecases.NewOrderUseCase(OrderRepository, ProductRepository)
	OrderHandler := handler.NewOrderHandler(OrderUseCase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	routes.CustomerRoutes(r, CustomerHandler)
	routes.ProductRoutes(r, ProductHandler)
	routes.OrderRoutes(r, OrderHandler)

	log.Println("API rodando em http://localhost:8080")
	log.Println("GET    /client      -> listar clientes")
	log.Println("POST	/client	->adicionar cliente")
	log.Println("DELETE	/client	->deletar cliente")
	log.Println("PUT	/client	->modificar cliente")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %s", err)
	}

}
