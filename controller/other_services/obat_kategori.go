package service

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/pa_user/helper"
	"github.com/hakushigo/pa_user/model/packets"
)

func ReqAddKategoriObat(ctx *fiber.Ctx) error {

	// recieve request
	var input packets.KategoriObat
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	// check the request
	if !(len(input.NamaKategoriObat) > 2) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  "The name needs to be more than 2 characters",
		})
	}

	// http://localhost:3001/obat/kategori
	// make a request agent to the product service
	agent := fiber.Post(helper.ObatServiceHostname + "/obat/kategori")
	agent.JSON(input) // forward the body to the service.

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

func ReqListKategoriObat(ctx *fiber.Ctx) error {
	agent := fiber.Get(helper.ObatServiceHostname + "/obat/kategori")
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

func ReqGetKategoriObat(ctx *fiber.Ctx) error {
	agent := fiber.Get(helper.ObatServiceHostname + "/obat/kategori/" + ctx.Params("id"))
	s, b, e := agent.Bytes()

	if len(e) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  e,
		})
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(s).Send(b)
}

func ReqUpdateKategoriObat(ctx *fiber.Ctx) error {
	var input packets.KategoriObat

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	// check the request
	if !(len(input.NamaKategoriObat) > 2) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  "The name is too short",
		})
	}

	id := ctx.Params("id")
	agent := fiber.Put(helper.ObatServiceHostname + "/obat/kategori/" + id)
	agent.JSON(input)

	s, b, e := agent.Bytes()
	if len(e) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  e,
		})
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(s).Send(b)
}

func ReqCountKategoriObat(ctx *fiber.Ctx) error {
	agent := fiber.Get(helper.ObatServiceHostname + "/obat/kategori")
	_, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  errs,
		})
	}

	var listOfKategori []packets.KategoriObat
	err := json.Unmarshal(body, &listOfKategori)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": 200,
		"count":  len(listOfKategori),
	})
}

func ReqDeleteKategoriObat(ctx *fiber.Ctx) error {
	agent := fiber.Delete(helper.ObatServiceHostname + "/obat/kategori/" + ctx.Params("id"))
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
