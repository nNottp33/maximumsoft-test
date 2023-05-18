package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nNottp33/maximumsoft-test/src/controllers"
	middleware "github.com/nNottp33/maximumsoft-test/src/middlewares"
)

func IndexRoute(app *fiber.App) {
	AuthRoute(app)

	common := app.Group("/api")
	common.Use(middleware.AuthMiddleware)
	EmployeeRoute(common)
}

func EmployeeRoute(router fiber.Router) {
	router.Get("/employees", controllers.GetEmployees)
	router.Post("/employee", controllers.NewEmployees)
	router.Put("/employee", controllers.UpdateEmployee)
	router.Delete("/employee/:id", controllers.RemoveEmployee)
}

func AuthRoute(router fiber.Router) {
	router.Post("/auth", controllers.SignIn)
}
