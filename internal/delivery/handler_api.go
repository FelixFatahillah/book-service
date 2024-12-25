package delivery

import (
	"book-service/internal/domain/book"
	"book-service/internal/domain/book/services"
	"book-service/pkg/exception"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type handler struct {
	serviceBook services.BookService
}

func NewHandler(serviceBook services.BookService) *handler {
	return &handler{
		serviceBook: serviceBook,
	}
}

const idleTimeout = 60 * time.Second

func (handler *handler) Init() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.FiberErrorHandler,
		IdleTimeout:  idleTimeout,
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(etag.New())
	app.Use(requestid.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	book.NewHandlerRESTBook(handler.serviceBook, app)

	return app
}
