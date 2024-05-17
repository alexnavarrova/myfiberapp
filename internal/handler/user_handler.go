package handler

import (
	"github.com/alexnavarrova/myfiberapp/internal/model"
	"github.com/alexnavarrova/myfiberapp/internal/service"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Hash the password before saving it
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not hash password",
        })
    }
    user.Password = string(hashedPassword)

    err = service.CreateUser(user)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not create user",
        })
    }

    return c.SendStatus(fiber.StatusNoContent)
}