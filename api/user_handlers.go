package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rojerdu-dev/hotel-reservation/types"
)

func HandleGetUsers(ctx fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Bond",
	}
	return ctx.JSON(u)
}

func HandleGetUser(ctx fiber.Ctx) error {
	return ctx.JSON("James")
}
