package main

import (
	"github.com/alexnavarrova/myfiberapp/config"
	"github.com/alexnavarrova/myfiberapp/internal/handler"
	"github.com/alexnavarrova/myfiberapp/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    
    config.LoadConfig()
    database.Connect()
    database.AutoMigrate()

    app.Get("/", handler.GetHome)
    app.Get("/users", handler.GetUsers)
    app.Post("/users", handler.CreateUser)

    app.Get("/books", handler.GetBooks)
    app.Post("/books", handler.CreateBook)

    app.Get("/libraries", handler.GetLibraries)
    app.Post("/libraries", handler.CreateLibrary)

    app.Listen(":3001")
}