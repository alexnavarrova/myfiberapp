package handler

import (
	"github.com/alexnavarrova/myfiberapp/internal/model"
	"github.com/alexnavarrova/myfiberapp/internal/repository"
	"github.com/alexnavarrova/myfiberapp/internal/util"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
    var req LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    var user model.User
    err := repository.FindUserByEmail(req.Email, &user)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid email or password",
        })
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid email or password",
        })
    }

    token, err := util.GenerateJWT(uint(user.ID))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not generate token",
        })
    }

    return c.JSON(fiber.Map{
        "token": token,
    })
}