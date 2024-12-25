package repositories

import (
	"book-service/internal/domain/book/models"
	"book-service/pkg/logger"
	"context"
)

func (repository repositoryBook) Loan(ctx context.Context, user models.BookLoaning) (*models.BookLoaning, error) {
	if err := repository.db.WithContext(ctx).Create(&user).Error; err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &user, nil
}

func (repository repositoryBook) Return(ctx context.Context, input *models.BookLoaning) error {
	if err := repository.db.
		WithContext(ctx).
		Updates(input).Error; err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (repository repositoryBook) FindLoanById(ctx context.Context, id string) (*models.BookLoaning, error) {
	record := &models.BookLoaning{}
	err := repository.db.WithContext(ctx).Take(record, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return record, nil
}
