package service

import (
	"github.com/alexnavarrova/myfiberapp/internal/model"
	"github.com/alexnavarrova/myfiberapp/internal/repository"
)

func GetAllUsers() ([]model.User, error) {
    return repository.GetAllUsers()
}

func CreateUser(user *model.User) error {
    return repository.CreateUser(user)
}
