package repositories

import (
	"book-service/internal/domain/book/dtos"
	"book-service/internal/domain/book/models"
	"book-service/pkg/exception"
	"book-service/pkg/helper"
	"book-service/pkg/logger"
	"context"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

func NewRepositoryBook(db *gorm.DB) *repositoryBook {
	return &repositoryBook{db}
}

func (repository repositoryBook) beginTransaction() *gorm.DB { return repository.db.Begin() }

func (repository repositoryBook) withTx(ctx context.Context, tx *gorm.DB) *repositoryBook {
	repository.db = tx.WithContext(ctx)
	return &repository
}

type ViewStockVersion struct {
	Stock   uint
	Version optimisticlock.Version
}

type ParamUpdateStock struct {
	BookID  string
	Version optimisticlock.Version
}

func (repository repositoryBook) Transaction(ctx context.Context, fn func(repo BookRepository) error) error {
	tx := repository.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	repo := repository.withTx(ctx, tx)
	err := fn(repo)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (repository repositoryBook) Create(ctx context.Context, user models.Book) (*models.Book, error) {
	if err := repository.db.WithContext(ctx).Create(&user).Error; err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &user, nil
}

func (repository repositoryBook) GetAll(ctx context.Context, filter dtos.BookFilter) ([]models.Book, *helper.PaginationMeta, error) {
	records := make([]models.Book, 0)
	paginateMeta := helper.PaginationMeta{
		Page:  filter.Pagination.Page,
		Limit: filter.Pagination.Limit,
	}

	query := repository.db.WithContext(ctx).Model(models.Book{}).Scopes(helper.PaginateScope(&filter.Pagination))
	query.Count(&paginateMeta.Total).Order("created_date DESC").Find(&records)

	paginateMeta.TotalPage = helper.GetTotalPage(paginateMeta.Total, paginateMeta.Limit)

	return records, &paginateMeta, nil
}

func (repository repositoryBook) FindById(ctx context.Context, id string) (*models.Book, error) {
	record := &models.Book{}
	err := repository.db.WithContext(ctx).Take(record, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return record, nil
}

func (repository repositoryBook) FindByTitle(ctx context.Context, title string) (*models.Book, error) {
	record := &models.Book{}
	err := repository.db.WithContext(ctx).Take(record, "title = ?", title).Error

	if err != nil {
		return nil, err
	}

	return record, nil
}

func (repository repositoryBook) FindStockById(ctx context.Context, id string) (*ViewStockVersion, error) {
	var result ViewStockVersion

	err := repository.db.
		WithContext(ctx).
		Model(models.Book{}).
		Select("stock, version").
		Take(&result, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repository repositoryBook) UpdateStock(ctx context.Context, stock uint, target ParamUpdateStock) error {
	trx := repository.db.
		WithContext(ctx).
		Model(&models.Book{
			ID:      target.BookID,
			Version: target.Version,
		})

	if err := trx.
		Update("stock", stock).Error; err != nil {
		logger.Error(err.Error())
		return err
	}

	if target.Version.Valid {
		if trx.RowsAffected == 0 {
			return exception.ErrOptimisticLock
		}
	}

	return nil
}

func (repository repositoryBook) Update(ctx context.Context, input *models.Book) error {
	if err := repository.db.
		WithContext(ctx).
		Updates(input).Error; err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (repository repositoryBook) Delete(ctx context.Context, id string) error {
	err := repository.db.WithContext(ctx).Unscoped().Where("id = ?", id).Delete(&models.Book{}).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
