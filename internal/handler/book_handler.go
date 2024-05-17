package handler

import (
	"github.com/alexnavarrova/myfiberapp/internal/model"
	"github.com/alexnavarrova/myfiberapp/internal/service"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
    books, err := service.GetAllBooks()
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(books)
}

func CreateBook(c *fiber.Ctx) error {
    book := new(model.Book)
    if err := c.BodyParser(book); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    err := service.CreateBook(book)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(book)
}