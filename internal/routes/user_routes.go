package routes

import (
	"github.com/alexnavarrova/myfiberapp/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
    app.Get("/users", handler.GetUsers)
    app.Post("/users", handler.CreateUser)
}