package routes

import (
	"github.com/gofiber/fiber/v2"
	"pesa_api/server/controllers"
)

func Routes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Group("/test/t1", controllers.Test)

}
