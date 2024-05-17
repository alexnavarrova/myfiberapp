package repository

import (
	"github.com/alexnavarrova/myfiberapp/internal/model"
	"github.com/alexnavarrova/myfiberapp/pkg/database"
)

func GetAllUsers() ([]model.User, error) {
    var users []model.User
    result := database.DB.Find(&users)
    return users, result.Error
}

func CreateUser(user *model.User) error {
    result := database.DB.Create(user)
    return result.Error
}
