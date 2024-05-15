package service

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/pa_user/helper"
	"github.com/hakushigo/pa_user/model/packets"
)

func ReqAddKategoriProduk(ctx *fiber.Ctx) error {
	// recieve an input of kategori
	var input packets.KategoriProduk
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

	agent := fiber.Post(helper.ProdukServiceHostname + "/produk/kategori")
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

func ReqListKategoriProduk(ctx *fiber.Ctx) error {
	agent := fiber.Get(helper.ProdukServiceHostname + "/produk/kategori")
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

func ReqGetKategoriProduk(ctx *fiber.Ctx) error {
	agent := fiber.Get(helper.ProdukServiceHostname + "/produk/kategori/" + ctx.Params("id"))
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

func ReqUpdateKategoriProduk(ctx *fiber.Ctx) error {
	var input packets.KategoriProduk

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
	agent := fiber.Put(helper.ProdukServiceHostname + "/produk/kategori/" + id)
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

func ReqCountKategoriProduk(ctx *fiber.Ctx) error {
	agent := fiber.Get(helper.ProdukServiceHostname + "/produk/kategori")

	_, body, errors := agent.Bytes()
	if len(errors) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  errors,
		})
	}

	var listOfKategoriProduk []packets.KategoriProduk
	err := json.Unmarshal(body, &listOfKategoriProduk)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": 200,
		"count":  len(listOfKategoriProduk),
	})
}

func ReqDeleteKategoriProduk(ctx *fiber.Ctx) error {
	agent := fiber.Delete(helper.ProdukServiceHostname + "/produk/kategori/" + ctx.Params("id"))
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
