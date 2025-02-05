package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/henrymoreirasilva/go-api/configs"
	"github.com/henrymoreirasilva/go-api/internal/entity"
	"github.com/henrymoreirasilva/go-api/internal/infra/database"
	"github.com/henrymoreirasilva/go-api/internal/infra/handlers"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)
	ProductHandler := handlers.NewHandlerProduct(productDB)

	userDB := database.NewUser(db)
	UserHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/products", ProductHandler.CreateProduct)
	r.Get("/products/", ProductHandler.GetProducts)
	r.Get("/products/{id}", ProductHandler.GetProduct)
	r.Put("/products/{id}", ProductHandler.UpdateProduct)
	r.Delete("/products/{id}", ProductHandler.DeleteProduct)

	r.Post("/users", UserHandler.CreateUser)

	http.ListenAndServe(":8000", r)
}
