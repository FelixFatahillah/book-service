package config

import (
	"book-service/internal/domain/book/models"
	loggerDefault "book-service/pkg/logger"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewClient() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		Viper().GetString("DB_HOST"),
		Viper().GetString("DB_USER"),
		Viper().GetString("DB_PASSWORD"),
		Viper().GetString("DB_NAME"),
		Viper().GetString("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&models.Book{}, &models.BookLoaning{}); err != nil {
		fmt.Println("Migration failed:", err)
		return nil, fmt.Errorf("failed to migrate models: %w", err)
	}
	fmt.Println("Migration successful!")

	if Viper().GetString("ENV") != "production" {
		db.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,       // Don't include params in the SQL log
				Colorful:                  true,        // Disable color
			},
		)
	}

	if err != nil {
		loggerDefault.Error("error migrating models", zap.Error(err))
		return nil, fmt.Errorf("failed to migrate models: %w", err)
	}

	return db, nil
}
