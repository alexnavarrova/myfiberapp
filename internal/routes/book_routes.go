package routes

import (
	"github.com/alexnavarrova/myfiberapp/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
    app.Get("/books", handler.GetBooks)
    app.Post("/books", handler.CreateBook)
}