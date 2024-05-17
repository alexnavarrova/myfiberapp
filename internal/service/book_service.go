package service

import (
	"github.com/alexnavarrova/myfiberapp/internal/model"
	"github.com/alexnavarrova/myfiberapp/internal/repository"
)

func GetAllBooks() ([]model.Book, error) {
    return repository.GetAllBooks()
}

func CreateBook(book *model.Book) error {
    return repository.CreateBook(book)
}