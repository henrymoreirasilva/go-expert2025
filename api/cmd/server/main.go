package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/henrymoreirasilva/go-api/configs"
	"github.com/henrymoreirasilva/go-api/internal/entity"
	"github.com/henrymoreirasilva/go-api/internal/infra/database"
	"github.com/henrymoreirasilva/go-api/internal/infra/handlers"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")
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
	UserHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", ProductHandler.CreateProduct)
		r.Get("/", ProductHandler.GetProducts)
		r.Get("/{id}", ProductHandler.GetProduct)
		r.Put("/{id}", ProductHandler.UpdateProduct)
		r.Delete("/{id}", ProductHandler.DeleteProduct)
	})

	r.Post("/users", UserHandler.CreateUser)
	r.Post("/users/generate_token", UserHandler.GetJWT)

	http.ListenAndServe(":8000", r)
}
