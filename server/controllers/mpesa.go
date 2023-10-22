package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/k0kubun/pp"
	"net/http"

	"pesa_api/package/model"
	"pesa_api/server"
)

func MpesaTransaction(c *fiber.Ctx) error {
	req := model.MpesaBroker{}
	if errDataBinding := c.XML(&req); errDataBinding != nil {
		log.Errorf("error binding data")
		log.Errorf("error occurred :%v\n", errDataBinding)
		return server.UnSuccessResponse(c, http.StatusInternalServerError, "invalid data submitted, try again")
	}
	pp.Println(req)
	return server.SuccessResponse(c, http.StatusOK, "ok")
}

func MpesaCallback(c *fiber.Ctx) error {
	req := model.MpesaBrokerResponseConfirming{}
	if errDataBinding := c.XML(&req); errDataBinding != nil {
		log.Errorf("error binding data ")
		log.Errorf("error occurred :%v\n", errDataBinding)
		return server.UnSuccessResponse(c, http.StatusInternalServerError, "invalid data submitted, try again")
	}
	pp.Println(req)
	return server.SuccessResponse(c, http.StatusOK, "ok")
}
