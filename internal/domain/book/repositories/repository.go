package repositories

import (
	"book-service/internal/domain/book/dtos"
	"book-service/internal/domain/book/models"
	"book-service/pkg/helper"
	"context"
	"gorm.io/gorm"
)

type repositoryBook struct {
	db *gorm.DB
}

type BookRepository interface {
	Transaction(ctx context.Context, fn func(repo BookRepository) error) error
	Create(ctx context.Context, input models.Book) (*models.Book, error)
	GetAll(ctx context.Context, filter dtos.BookFilter) ([]models.Book, *helper.PaginationMeta, error)
	FindById(ctx context.Context, id string) (*models.Book, error)
	FindByTitle(ctx context.Context, title string) (*models.Book, error)
	FindStockById(ctx context.Context, id string) (*ViewStockVersion, error)
	UpdateStock(ctx context.Context, stock uint, target ParamUpdateStock) error
	Update(ctx context.Context, input *models.Book) error
	Delete(ctx context.Context, id string) error

	Loan(ctx context.Context, input models.BookLoaning) (*models.BookLoaning, error)
	Return(ctx context.Context, input *models.BookLoaning) error
	FindLoanById(ctx context.Context, id string) (*models.BookLoaning, error)
}
