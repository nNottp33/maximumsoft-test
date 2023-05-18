package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nNottp33/maximumsoft-test/src/controllers"
)

func IndexRoute(app *fiber.App) {
	EmployeeRoute(app)
}

func EmployeeRoute(router fiber.Router) {
	router.Get("/employees", controllers.GetEmployees)
}
