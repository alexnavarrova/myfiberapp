package routes

import (
	"github.com/alexnavarrova/myfiberapp/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func LibraryRoutes(app *fiber.App) {
    app.Get("/libraries", handler.GetLibraries)
    app.Post("/libraries", handler.CreateLibrary)
}