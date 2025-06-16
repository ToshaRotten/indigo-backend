package main

import (
	"fmt"

	"main/db/wrapper"
)

func main() {
	// TODO: create a config service with config file
	pgDSN := "host=172.17.0.2 port=5432 user=postgres password=pass dbname=postgres sslmode=disable"
	db, err := wrapper.OpenConnection(pgDSN)
	if err != nil {
		fmt.Println(err)
	}

	db.CloseConnection()

	//domain.Initialize(db)
	//
	// app := fiber.New()
	//
	// // TODO: use a normal middleware
	// app.Use(func(c fiber.Ctx) error {
	// 	fmt.Println(c.Request())
	// 	return c.Next()
	// })
	//
	// api := app.Group("/api")
	//
	// // Service
	// api.Get("/status", status.Status)
	//
	// // User
	// userController := user.NewUserController()
	//
	// api.Post("/user", userController.Create)
	// api.Delete("/user", userController.Remove)
	// api.Get("/user/:id", userController.Get)
	// api.Put("/user", userController.Update)
	//
	// // Service auth
	// api.Post("/user/login", auth.Login)
	// api.Post("/user/register", auth.Registration)
	//
	// // Project
	// project := project.NewProjectController()
	//
	// api.Post("/project", project.Create)
	// api.Get("/project/:id", project.Get)
	// api.Put("/project", project.Update)
	// api.Delete("/project/:id", project.Remove)
	//
	// log.Fatal(app.Listen(":3000"))
}
