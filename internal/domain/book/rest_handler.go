package book

import (
	"book-service/internal/domain/book/services"
	"book-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type handlerRESTBook struct {
	service services.BookService
}

func NewHandlerRESTBook(service services.BookService, router *fiber.App) {
	handler := handlerRESTBook{
		service,
	}

	routerV1 := router.Group("/api/v1")
	routerProtected := routerV1.Group("/private/books", middleware.RoleMiddleware())

	routerProtected.Post("/loan", handler.handlerLoan)
	routerProtected.Post("/return", handler.handlerReturn)

	routerProtected.Delete("/:id", handler.handlerDelete)
	routerProtected.Get("/:id", handler.handlerFindById)
	routerProtected.Put("/:id", handler.handlerUpdate)
	routerProtected.Post("/", handler.handlerCreate)
	routerProtected.Get("/", handler.handlerGetAll)
}
