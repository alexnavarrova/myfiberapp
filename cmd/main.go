package main

import (
	"github.com/alexnavarrova/myfiberapp/config"
	"github.com/alexnavarrova/myfiberapp/internal/routes"
	"github.com/alexnavarrova/myfiberapp/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    
    config.LoadConfig()
    database.Connect()
    database.AutoMigrate()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // Registrar rutas
    routes.UserRoutes(app)
    routes.BookRoutes(app)
    routes.LibraryRoutes(app)

    app.Listen(":3001")
}