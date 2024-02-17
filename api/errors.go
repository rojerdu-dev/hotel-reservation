package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(Error); ok {
		return c.Status(apiError.Code).JSON(apiError)
	}
	apiError := NewError(http.StatusInternalServerError, err.Error())
	return c.Status(apiError.Code).JSON(apiError)
}

type Error struct {
	Code   int    `json:"code"`
	ErrMsg string `json:"error"`
}

// Error implements the Error interface
func (e Error) Error() string {
	return e.ErrMsg
}

func NewError(code int, err string) Error {
	return Error{
		Code:   code,
		ErrMsg: err,
	}
}

func ErrorResourceNotFound(res string) Error {
	return Error{
		Code:   http.StatusNotFound,
		ErrMsg: res + "resource not found",
	}
}

func ErrorBadRequest() Error {
	return Error{
		Code:   http.StatusBadRequest,
		ErrMsg: "invalid JSON request",
	}
}

func ErrorUnauthorized() Error {
	return Error{
		Code:   http.StatusUnauthorized,
		ErrMsg: "unauthorized request",
	}
}

func ErrInvalidID() Error {
	return Error{
		Code:   http.StatusBadRequest,
		ErrMsg: "invalid id given",
	}
}
