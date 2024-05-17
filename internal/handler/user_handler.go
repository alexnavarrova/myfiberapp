package handler

import (
	"github.com/alexnavarrova/myfiberapp/internal/model"
	"github.com/alexnavarrova/myfiberapp/internal/service"
	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
}

func GetUsers(c *fiber.Ctx) error {
    users, err := service.GetAllUsers()
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
    user := new(model.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    err := service.CreateUser(user)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(user)
}