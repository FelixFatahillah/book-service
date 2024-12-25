package exception

import (
	"book-service/pkg/response"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

func FiberErrorHandler(ctx *fiber.Ctx, err error) error {
	defaultRes := response.HTTPResponse{
		Code:    fiber.StatusInternalServerError,
		Message: "exception.internal_error",
	}

	var errValidation *ErrValidation
	if errors.As(err, &errValidation) {
		data := errValidation.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicIfNeeded(errJson)

		defaultRes.Code = fiber.StatusBadRequest
		defaultRes.Message = "exception.bad_request"
		var errs []interface{}
		for _, message := range messages {
			errs = append(errs, message)
		}
		defaultRes.Errors = errs
	}

	var withCodeErr *ErrWithCode
	if errors.As(err, &withCodeErr) {
		defaultRes.Code = http.StatusInternalServerError
		if withCodeErr.Code > 0 {
			defaultRes.Code = withCodeErr.Code
		}
		defaultRes.Message = http.StatusText(defaultRes.Code)
		if withCodeErr.Err != nil {
			defaultRes.Message = withCodeErr.Err.Error()
		}
	}

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		defaultRes.Code = fiberError.Code
		defaultRes.Message = fiberError.Message
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		defaultRes.Code = fiber.StatusNotFound
		defaultRes.Message = "exception.not_found"
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		defaultRes.Code = fiber.StatusConflict
		defaultRes.Message = "exception.already_exist"
	}

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		defaultRes.Code = fiber.StatusUnprocessableEntity
		defaultRes.Message = http.StatusText(fiber.StatusUnprocessableEntity)

		defaultRes.Errors = []interface{}{
			map[string]interface{}{
				"field":   unmarshalTypeError.Field,
				"message": fmt.Sprintf("%s must be %s", unmarshalTypeError.Field, unmarshalTypeError.Type),
			},
		}
	}

	if defaultRes.Code >= 500 {
		defaultRes.Message = "exception.unknown_error"
	}

	return ctx.Status(defaultRes.Code).JSON(defaultRes)
}

func PanicIfNeeded(err error) {
	if err != nil {
		panic(err)
	}
}
