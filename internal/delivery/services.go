package delivery

import (
	servicesBook "book-service/internal/domain/book/services"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	BookService servicesBook.BookService
}

type Deps struct {
	Repository *Repositories
	Redis      redis.Client
	GRPC       *GRPC
}

func NewService(deps Deps) *Service {
	return &Service{
		BookService: servicesBook.NewServiceBook(deps.Repository.BookRepository, deps.Redis, deps.GRPC.AuthorRPC, deps.GRPC.CategoryRPC),
	}
}
