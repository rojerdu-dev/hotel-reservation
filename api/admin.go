package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rojerdu-dev/hotel-reservation/types"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return ErrorUnauthorized()
		//return fmt.Errorf("not authorized")
	}
	if !user.IsAdmin {
		return ErrorUnauthorized()
		//return fmt.Errorf("not authorized")
	}
	return c.Next()
}
