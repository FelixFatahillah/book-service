package middleware

import (
	"book-service/pkg/exception"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	_ "github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func RateLimit(max int, expiration time.Duration, message string) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: expiration,
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.Get("X-Forwarded-For")
		},
		LimitReached: func(_ *fiber.Ctx) error {
			return &exception.ErrWithCode{
				Code: fiber.StatusTooManyRequests,
				Err:  errors.New(message),
			}
		},
	})
}
