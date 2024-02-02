package main

import "github.com/gofiber/fiber/v3"

func main() {
	//fmt.Println("starting hotel-reservation project")

	app := fiber.New()
	app.Get("/home", home)
	app.Listen(":5000")

}

func home(ctx fiber.Ctx) error {
	return ctx.JSON(map[string]string{"msg": "home is where the heart is"})
}
