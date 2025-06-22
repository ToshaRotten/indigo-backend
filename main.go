package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"

	"main/db/seed"
	"main/db/wrapper"
	"main/domain"
	"main/handlers/auth"
	"main/handlers/middleware"
	"main/handlers/product"
	"main/handlers/status"
	"main/handlers/user"
)

func main() {
	// TODO: create a config service with config file
	pgDSN := "host=172.17.0.2 port=5432 user=postgres password=pass dbname=postgres sslmode=disable"
	db, err := wrapper.OpenConnection(pgDSN)
	if err != nil {
		fmt.Println(err)
	}

	domain.Initialize(db)
	seed.SeedDatabase(context.Background(), db.Client)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// TODO: use a normal middleware
	app.Use(func(c fiber.Ctx) error {
		fmt.Println(c.Request())
		return c.Next()
	})

	domain.Initialize(db)

	api := app.Group("/api")
	app.Get("/images*", static.New("", static.Config{
		FS:     os.DirFS("uploads/images"),
		Browse: true,
	}))

	// Service
	api.Get("/status", status.Status)

	authApi := app.Group("/auth")

	// Service auth
	authApi.Post("/user/login", auth.Login)
	authApi.Post("/user/register", auth.Registration)
	api.Use(middleware.JwtMiddleware)

	user := user.NewUserController()
	api.Get("/users", user.GetAll)
	//
	// Project
	product := product.NewProductController()

	api.Post("/product", product.Create)
	api.Get("/product/:id", product.Get)
	api.Put("/product", product.Update)
	api.Delete("/product/:id", product.Remove)
	api.Get("/products", product.GetAll)

	log.Fatal(app.Listen(":3000"))
}
