package main

import (
	"flag"

	"github.com/gofiber/fiber/v3"
	"github.com/rojerdu-dev/hotel-reservation/api"
)

func main() {
	addr := flag.String("address", ":5000", "API server listen address")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/home", handleHome)
	apiv1.Get("/user/:id", api.HandleGetUser)
	apiv1.Get("/user", api.HandleGetUsers)
	app.Listen(*addr)
}

func handleHome(ctx fiber.Ctx) error {
	return ctx.JSON(map[string]string{"msg": "home is where the heart is"})
}
