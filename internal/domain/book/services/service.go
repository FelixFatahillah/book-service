package services

import (
	"book-service/internal/domain/book/dtos"
	"book-service/internal/domain/book/models"
	"book-service/internal/domain/book/repositories"
	"book-service/internal/infrastructure/pb"
	"book-service/pkg/helper"
	"context"
	"github.com/redis/go-redis/v9"
)

type serviceBook struct {
	Repository  repositories.BookRepository
	Redis       redis.Client
	RPCAuthor   pb.AuthorServiceClient
	RPCCategory pb.CategoryServiceClient
}

type BookService interface {
	Create(ctx context.Context, input dtos.CreateBookDto) (*dtos.CreateBookResponseDto, error)
	GetAll(ctx context.Context, filter dtos.BookFilter) ([]models.Book, *helper.PaginationMeta, error)
	FindById(ctx context.Context, id string) (*models.Book, error)
	Update(ctx context.Context, input dtos.UpdateBookDto) error
	Delete(ctx context.Context, id string) error

	Loan(ctx context.Context, input dtos.LoanDto) (*dtos.LoanDto, error)
	Return(ctx context.Context, input dtos.ReturnDto) (*dtos.ReturnDto, error)
}
