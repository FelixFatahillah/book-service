package services

import (
	"book-service/internal/domain/book/dtos"
	"book-service/internal/domain/book/models"
	"book-service/internal/domain/book/repositories"
	"book-service/internal/infrastructure/pb"
	"book-service/internal/shared"
	"book-service/pkg/exception"
	"book-service/pkg/helper"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

func NewServiceBook(repository repositories.BookRepository, redis redis.Client, authorRPC pb.AuthorServiceClient, categoryRPC pb.CategoryServiceClient) *serviceBook {
	return &serviceBook{
		Repository:  repository,
		Redis:       redis,
		RPCAuthor:   authorRPC,
		RPCCategory: categoryRPC,
	}
}

func (service serviceBook) Create(ctx context.Context, input dtos.CreateBookDto) (*dtos.CreateBookResponseDto, error) {
	categoryExist, _ := service.Repository.FindByTitle(ctx, input.Title)
	if categoryExist != nil {
		return nil, &exception.ErrWithCode{
			Code: http.StatusBadRequest,
			Err:  errors.New("book is already registered"),
		}
	}

	author, err := service.RPCAuthor.GetAuthorById(ctx, &pb.GetAuthorByIdRequest{Id: input.AuthorID})
	if err != nil {
		return nil, err
	}

	category, err := service.RPCCategory.GetCategoryById(ctx, &pb.GetCategoryByIdRequest{Id: input.CategoryID})
	if err != nil {
		return nil, err
	}

	record, err := service.Repository.Create(ctx, models.Book{
		Title:     input.Title,
		Genre:     input.Genre,
		Stock:     input.Stock,
		Published: input.Published,
		Author: models.Author{
			FirstName:   author.FirstName,
			LastName:    &author.LastName,
			PhoneNumber: &author.PhoneNumber,
			Email:       author.Email,
		},
		Category: models.Category{
			Name:        category.Category,
			Description: category.Description,
		},
	})
	if err != nil {
		return nil, err
	}
	return &dtos.CreateBookResponseDto{
		Title:     record.Title,
		Genre:     record.Genre,
		Stock:     record.Stock,
		Published: record.Published,
		Author: models.Author{
			FirstName:   record.Author.FirstName,
			LastName:    record.Author.LastName,
			PhoneNumber: record.Author.PhoneNumber,
			Email:       record.Author.Email,
		},
		Category: models.Category{
			Name:        record.Category.Name,
			Description: record.Category.Description,
		},
	}, nil
}
func (service serviceBook) GetAll(ctx context.Context, filter dtos.BookFilter) ([]models.Book, *helper.PaginationMeta, error) {
	return service.Repository.GetAll(ctx, filter)
}

func (service serviceBook) FindById(ctx context.Context, id string) (*models.Book, error) {
	return service.Repository.FindById(ctx, id)
}

func (service serviceBook) Update(ctx context.Context, input dtos.UpdateBookDto) error {
	err := service.Repository.Transaction(ctx, func(repo repositories.BookRepository) error {
		record, _ := service.Repository.FindById(ctx, input.ID)
		if record == nil {
			return &exception.ErrWithCode{
				Code: http.StatusNotFound,
				Err:  errors.New("category not found"),
			}
		}
		_ = service.Repository.Update(ctx,
			&models.Book{
				ID: record.ID,
				BaseModel: shared.BaseModel{
					UpdatedDate: time.Now().Local(),
				},
				Title:     input.Title,
				Genre:     input.Genre,
				Stock:     input.Stock,
				Published: input.Published,
				Author: models.Author{
					FirstName:   input.Author.FirstName,
					LastName:    input.Author.LastName,
					PhoneNumber: input.Author.PhoneNumber,
					Email:       input.Author.Email,
				},
				Category: models.Category{
					Name:        input.Category.Name,
					Description: input.Category.Description,
				},
			},
		)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (service serviceBook) Delete(ctx context.Context, id string) error {
	category, _ := service.Repository.FindById(ctx, id)
	if category == nil {
		return &exception.ErrWithCode{
			Code: http.StatusNotFound,
			Err:  errors.New("category not found"),
		}
	}
	return service.Repository.Delete(ctx, id)
}
