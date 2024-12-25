package services

import (
	"book-service/internal/domain/book/dtos"
	"book-service/internal/domain/book/models"
	"book-service/internal/domain/book/repositories"
	"book-service/pkg/exception"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (service serviceBook) Loan(ctx context.Context, input dtos.LoanDto) (*dtos.LoanDto, error) {
	book, err := service.Repository.FindById(ctx, input.BookID)
	if err != nil {
		return nil, &exception.ErrWithCode{
			Code: http.StatusNotFound,
			Err:  errors.New("book not found"),
		}
	}

	maxRetries := 5
	for retries := 0; retries < maxRetries; retries++ {
		data, err := service.Repository.FindStockById(ctx, input.BookID)
		if err != nil {
			return nil, err
		}

		if data.Stock <= 0 {
			return nil, &exception.ErrWithCode{
				Code: http.StatusBadRequest,
				Err:  errors.New("book out of stock"),
			}
		}
		fmt.Println("stock: ", data.Stock)
		err = service.Repository.UpdateStock(ctx, data.Stock-1, repositories.ParamUpdateStock{
			BookID:  input.BookID,
			Version: data.Version,
		})
		fmt.Println("retriess: ", retries)
		if err == nil {
			break
		}

		if errors.Is(err, exception.ErrOptimisticLock) {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		return nil, err
	}

	if err != nil {
		return nil, &exception.ErrWithCode{
			Code: http.StatusConflict,
			Err:  errors.New("could not update stock due to concurrent operations"),
		}
	}

	record, err := service.Repository.Loan(ctx, models.BookLoaning{
		CustomerName:       input.CustomerName,
		LoanDate:           input.LoanDate,
		ReturnDateSchedule: input.ReturnDateSchedule,
		BookID:             input.BookID,
		BookTitle:          book.Title,
	})
	if err != nil {
		return nil, err
	}

	return &dtos.LoanDto{
		CustomerName:       record.CustomerName,
		LoanDate:           record.LoanDate,
		ReturnDateSchedule: record.ReturnDateSchedule,
		BookID:             record.BookID,
	}, nil
}

func (service serviceBook) Return(ctx context.Context, input dtos.ReturnDto) (*dtos.ReturnDto, error) {
	record, err := service.Repository.FindLoanById(ctx, input.LoanID)
	if err != nil {
		return nil, &exception.ErrWithCode{
			Code: http.StatusNotFound,
			Err:  errors.New("loan record not found"),
		}
	}

	if record.ReturnDate != nil {
		return nil, &exception.ErrWithCode{
			Code: http.StatusBadRequest,
			Err:  errors.New("loan transaction is already completed"),
		}
	}

	maxRetries := 5
	for retries := 0; retries < maxRetries; retries++ {
		data, err := service.Repository.FindStockById(ctx, record.BookID)
		if err != nil {
			return nil, err
		}

		err = service.Repository.UpdateStock(ctx, data.Stock+1, repositories.ParamUpdateStock{
			BookID:  record.BookID,
			Version: data.Version,
		})

		if err == nil {
			break
		}

		if errors.Is(err, exception.ErrOptimisticLock) {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		return nil, err
	}

	if err != nil {
		return nil, &exception.ErrWithCode{
			Code: http.StatusConflict,
			Err:  errors.New("could not update stock due to concurrent operations"),
		}
	}

	currentTime := time.Now().Local()
	err = service.Repository.Return(ctx, &models.BookLoaning{
		ID:         input.LoanID,
		ReturnDate: &currentTime,
	})
	if err != nil {
		return nil, err
	}

	return &dtos.ReturnDto{
		LoanID: record.ID,
	}, nil
}
