package routes

import (
	"github.com/gofiber/fiber/v2"
	"pesa_api/server/middlewares"
)

func Routers(app *fiber.App) {
	middlewares.Logger(app)
	Routes(app)
}
