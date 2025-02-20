// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Henry Silva
// @contact.url    http://www.henrymoreirasilva.com.br
// @contact.email  henry@zoomagencia.com.br

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/henrymoreirasilva/go-api/configs"
	_ "github.com/henrymoreirasilva/go-api/docs"
	"github.com/henrymoreirasilva/go-api/internal/entity"
	"github.com/henrymoreirasilva/go-api/internal/infra/database"
	"github.com/henrymoreirasilva/go-api/internal/infra/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
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

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}
