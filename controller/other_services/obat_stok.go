package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/pa_user/helper"
)

func AddStokObat(ctx *fiber.Ctx) error {
	// creating an agent
	agent := fiber.Put(helper.StokObatServiceHostname + "/obat/stok/add")
	agent.Body(ctx.Body())

	status, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  errs,
		})
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(status).Send(body)
}

func ReduceStokObat(ctx *fiber.Ctx) error {
	// creating an agent
	agent := fiber.Put(helper.StokObatServiceHostname + "/obat/stok/reduce")
	agent.Body(ctx.Body())

	status, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  errs,
		})
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(status).Send(body)
}
