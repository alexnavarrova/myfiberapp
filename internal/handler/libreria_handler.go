package handler

import (
	"github.com/alexnavarrova/myfiberapp/internal/model"
	"github.com/alexnavarrova/myfiberapp/internal/service"
	"github.com/gofiber/fiber/v2"
)

func GetLibraries(c *fiber.Ctx) error {
    libraries, err := service.GetAllLibraries()
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(libraries)
}

func CreateLibrary(c *fiber.Ctx) error {
    library := new(model.Library)
    if err := c.BodyParser(library); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    err := service.CreateLibrary(library)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(library)
}