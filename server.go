package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	config "github.com/nNottp33/maximumsoft-test/src/configs"
	"github.com/nNottp33/maximumsoft-test/src/routes"
)

func main() {
	app := fiber.New()

	// setup logger
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))
	// setup cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	// get routes
	routes.IndexRoute(app)

	err := app.Listen("localhost:" + config.PORT)
	if err != nil {
		panic(err)
	}
}
