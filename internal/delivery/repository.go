package delivery

import (
	repositoryBook "book-service/internal/domain/book/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	BookRepository repositoryBook.BookRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		BookRepository: repositoryBook.NewRepositoryBook(db),
	}
}
