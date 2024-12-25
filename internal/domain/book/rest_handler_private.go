package book

import (
	"book-service/internal/domain/book/dtos"
	"book-service/pkg/helper"
	"book-service/pkg/response"
	"book-service/pkg/validation"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (handler handlerRESTBook) handlerCreate(ctx *fiber.Ctx) error {
	var body dtos.CreateBookDto
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	if err := validation.Validate(body); err != nil {
		return err
	}

	data, err := handler.service.Create(ctx.Context(), body)
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, "success", data, nil, nil)
}

func (handler handlerRESTBook) handlerGetAll(ctx *fiber.Ctx) error {
	paginate := helper.Pagination{
		Page:  1,
		Limit: 10,
	}
	err := ctx.QueryParser(&paginate)
	if paginate.Limit >= 100 {
		paginate.Limit = 100
	}

	filter := dtos.BookFilter{
		Pagination: paginate,
	}
	_ = ctx.QueryParser(&filter)

	data, meta, err := handler.service.GetAll(ctx.Context(), filter)
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, "success", data, nil, meta)
}

func (handler handlerRESTBook) handlerFindById(ctx *fiber.Ctx) error {
	data, err := handler.service.FindById(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, "success", data, nil, nil)
}

func (handler handlerRESTBook) handlerUpdate(ctx *fiber.Ctx) error {
	var body dtos.UpdateBookDto
	body.ID = ctx.Params("id")
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	if err := validation.Validate(body); err != nil {
		return err
	}

	err := handler.service.Update(ctx.Context(), body)
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, fmt.Sprintf("success update %s", ctx.Params("id")), nil, nil, nil)
}

func (handler handlerRESTBook) handlerDelete(ctx *fiber.Ctx) error {
	err := handler.service.Delete(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, fmt.Sprintf("success deleted %s", ctx.Params("id")), nil, nil, nil)
}

func (handler handlerRESTBook) handlerLoan(ctx *fiber.Ctx) error {
	var body dtos.LoanDto
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	if err := validation.Validate(body); err != nil {
		return err
	}

	data, err := handler.service.Loan(ctx.Context(), body)
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, "success", data, nil, nil)
}

func (handler handlerRESTBook) handlerReturn(ctx *fiber.Ctx) error {
	var body dtos.ReturnDto
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	if err := validation.Validate(body); err != nil {
		return err
	}

	data, err := handler.service.Return(ctx.Context(), body)
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, "success", data, nil, nil)
}
