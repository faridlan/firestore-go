package exception

import (
	"github.com/faridlan/firestore-go/model/web"
	"github.com/gofiber/fiber/v2"
)

func ExceptionError(ctx *fiber.Ctx, err error) error {

	switch e := err.(type) {
	case *NotFound:
		return notFoundError(ctx, e.Error())
	case *BadRequest:
		return badRequestError(ctx, e.Error())
	default:
		return internalServerError(ctx, err)
	}

}

func internalServerError(ctx *fiber.Ctx, err error) error {

	ctx.Request().Header.Add("content-type", "application/json")
	ctx.Status(fiber.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err.Error(),
	}

	return ctx.JSON(webResponse)
}

func badRequestError(ctx *fiber.Ctx, err string) error {

	ctx.Request().Header.Add("content-type", "application/json")
	ctx.Status(fiber.StatusBadRequest)

	webResponse := web.WebResponse{
		Code:   fiber.StatusBadRequest,
		Status: "BAD REQUEST",
		Data:   err,
	}

	return ctx.JSON(webResponse)

}

func notFoundError(ctx *fiber.Ctx, err string) error {

	ctx.Request().Header.Add("content-type", "application/json")
	ctx.Status(fiber.StatusNotFound)

	webResponse := web.WebResponse{
		Code:   fiber.StatusNotFound,
		Status: "NOT FOUND",
		Data:   err,
	}

	return ctx.JSON(webResponse)

}
