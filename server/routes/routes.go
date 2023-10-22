package routes

import (
	"github.com/gofiber/fiber/v2"
	"pesa_api/server/controllers"
)

func Routes(app *fiber.App) {
	api := app.Group("/api/v0")
	api.Group("/test/t1", controllers.Test)

	vodacom := api.Group("/vodacom")
	{
		vodacom.Post("/transaction", controllers.MpesaTransaction)
		vodacom.Post("/callback", controllers.MpesaCallback)
	}

}
