package delivery

import (
	"book-service/internal/infrastructure/pb"
	"google.golang.org/grpc"
)

type GRPC struct {
	AuthorRPC   pb.AuthorServiceClient
	CategoryRPC pb.CategoryServiceClient
}

func NewGRPC(authorConn, categoryConn grpc.ClientConnInterface) *GRPC {
	return &GRPC{
		AuthorRPC:   pb.NewAuthorServiceClient(authorConn),
		CategoryRPC: pb.NewCategoryServiceClient(categoryConn),
	}
}
