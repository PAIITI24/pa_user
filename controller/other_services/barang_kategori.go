package service

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/pa_user/helper"
	"github.com/hakushigo/pa_user/model/packets"
	"strconv"
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

	if !(len(input.NamaKategoriBarang) > 2) {
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
	var input struct {
		NamaKategoriBarang string `json:"nama_kategori_barang"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	// check the request
	if !(len(input.NamaKategoriBarang) > 2) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  "The name is too short",
		})
	}

	id, _ := ctx.ParamsInt("id")
	agent := fiber.Put(helper.BarangServiceHostname + "/barang/kategori/" + strconv.Itoa(id))
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
