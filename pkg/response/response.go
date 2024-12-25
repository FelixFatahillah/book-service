package response

import (
	"github.com/gofiber/fiber/v2"
)

type HTTPResponse struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message,omitempty" example:"Success"`
	Data    any         `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Errors  []any       `json:"errors,omitempty" `
}

func Respond(c *fiber.Ctx, statusCode int, message string, data interface{}, err error, meta interface{}) error {
	response := HTTPResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
		Meta:    meta,
	}

	return c.Status(statusCode).JSON(response)
}
