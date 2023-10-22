package server

import "github.com/gofiber/fiber/v2"

type SuccessResponseStrct struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
}
type ErrorResponse struct {
	Message string `json:"error"`
}

func SuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.JSON(SuccessResponseStrct{
		Code: statusCode,
		Data: data,
	})
}
