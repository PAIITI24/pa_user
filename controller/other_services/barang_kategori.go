package service

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/pa_user/helper"
	"github.com/hakushigo/pa_user/model/packets"
)

func ReqAddKategoriBarang(ctx *fiber.Ctx) error {
	// recieve an input of kategori
	var input packets.KategoriBarang
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	if !(len(input.NamaKategori) > 2) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  "The name needs to be more than 2 characters",
		})
	}

	agent := fiber.Post(helper.BarangServiceHostname + "/barang/kategori")
	agent.JSON(input)

	status, body, errors := agent.Bytes()

	if len(errors) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  errors,
		})
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(status).Send(body)
}

func ReqListKategoriBarang(ctx *fiber.Ctx) error {
	agent := fiber.Get(helper.BarangServiceHostname + "/barang/kategori")
	status, body, errors := agent.Bytes()

	if len(errors) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  errors,
		})
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(status).Send(body)
}

func ReqGetKategoriBarang(ctx *fiber.Ctx) error {
	agent := fiber.Get(helper.BarangServiceHostname + "/barang/kategori/" + ctx.Params("id"))
	status, body, errors := agent.Bytes()

	if len(errors) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  errors,
		})
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(status).Send(body)
}

func ReqUpdateKategoriBarang(ctx *fiber.Ctx) error {
	var input packets.KategoriBarang

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	if !(len(input.NamaKategori) > 2) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  "The name needs to be more than 2 characters",
		})
	}

	id := ctx.Params("id")
	agent := fiber.Put(helper.BarangServiceHostname + "/barang/kategori/" + id)
	agent.JSON(input)

	status, body, errors := agent.Bytes()
	if len(errors) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  errors,
		})
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(status).Send(body)
}

func ReqCountKategoriBarang(ctx *fiber.Ctx) error {
	agent := fiber.Get(helper.BarangServiceHostname + "/barang/kategori")

	_, body, errors := agent.Bytes()
	if len(errors) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  errors,
		})
	}

	var listOfKategoriBarang []packets.KategoriBarang
	err := json.Unmarshal(body, &listOfKategoriBarang)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": 200,
		"count":  len(listOfKategoriBarang),
	})
}

func ReqDeleteKategoriBarang(ctx *fiber.Ctx) error {
	agent := fiber.Delete(helper.BarangServiceHostname + "/barang/kategori/" + ctx.Params("id"))
	status, body, errors := agent.Bytes()

	if len(errors) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  errors,
		})
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(status).Send(body)
}
