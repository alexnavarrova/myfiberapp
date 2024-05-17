package repository

import (
	"github.com/alexnavarrova/myfiberapp/internal/model"
	"github.com/alexnavarrova/myfiberapp/pkg/database"
)

func GetAllBooks() ([]model.Book, error) {
    var books []model.Book
    result := database.DB.Find(&books)
    return books, result.Error
}

func CreateBook(book *model.Book) error {
    result := database.DB.Create(book)
    return result.Error
}
