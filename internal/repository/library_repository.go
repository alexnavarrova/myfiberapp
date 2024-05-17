package repository

import (
	"github.com/alexnavarrova/myfiberapp/internal/model"
	"github.com/alexnavarrova/myfiberapp/pkg/database"
)

func GetAllLibraries() ([]model.Library, error) {
    var libraries []model.Library
    result := database.DB.Find(&libraries)
    return libraries, result.Error
}

func CreateLibrary(library *model.Library) error {
    result := database.DB.Create(library)
    return result.Error
}
