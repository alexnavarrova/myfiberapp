package service

import (
	"github.com/alexnavarrova/myfiberapp/internal/model"
	"github.com/alexnavarrova/myfiberapp/internal/repository"
)

func GetAllLibraries() ([]model.Library, error) {
    return repository.GetAllLibraries()
}

func CreateLibrary(library *model.Library) error {
    return repository.CreateLibrary(library)
}