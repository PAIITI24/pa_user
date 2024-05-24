package service

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/pa_user/helper"
	"github.com/hakushigo/pa_user/model/packets"
	"github.com/minio/minio-go/v7"
	"path/filepath"
)

func ReqAddBarang(ctx *fiber.Ctx) error {
	//	var newBarang packets.Barang
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"statusn": 500,
			"error":   err.Error(),
		})
	}

	// extract data
	var data struct {
		KategoriBarang []int          `json:"kategori_barang"`
		DataBarang     packets.Barang `json:"data_barang"`
	}

	if len(form.Value["data"]) == 0 || len(form.File["image"]) == 0 { // check if both data and image is here
		return ctx.Status(500).JSON(fiber.Map{
			"statusn": 500,
			"error":   "Both photo and file object should be in the request",
		})
	}

	err = json.Unmarshal([]byte(form.Value["data"][0]), &data) // decode to data struct object
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"statusn": 500,
			"error":   err.Error(),
		})
	}

	// extract uploaded file
	image := form.File["image"][0]
	imageContent, err := image.Open()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"statusn": 500,
			"error":   err.Error(),
		})
	}

	//	imgBBuffer := bytes.NewBuffer(nil) // create a buffer for the image
	objectName := helper.GenerateToken(image.Filename) + filepath.Ext(image.Filename)

	// preparing a client
	agent := fiber.Post(helper.BarangServiceHostname + "/barang/")

	data.DataBarang.Gambar = helper.BarangStoragePublicURL + "/" + objectName // adding gambar URL before sending it
	agent.JSON(data)                                                          // push push push

	s, b, e := agent.Bytes()

	if len(e) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  e,
		})
	}

	if s == 200 { // only upload file to the S3 if the data to the database is already inserted

		S3Client := helper.S3Connect()
		uploadCTX := context.Background() //S3Context

		_, err = S3Client.PutObject(uploadCTX, helper.BarangBucketName, objectName, imageContent, image.Size, minio.PutObjectOptions{
			ContentType: image.Header.Get("Content-Type"),
		})
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"statusn": 500,
				"error":   "failed to upload data, an error occured : " + err.Error(),
			})
		}
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(s).Send(b)
}

func ReqListBarang(ctx *fiber.Ctx) error {
	// creating an agent
	agent := fiber.Get(helper.BarangServiceHostname + "/barang")
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

func ReqGetBarang(ctx *fiber.Ctx) error {
	// creating an agent
	agent := fiber.Get(helper.BarangServiceHostname + "/barang/" + ctx.Params("id"))
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

func ReqUpdateBarang(ctx *fiber.Ctx) error {
	//	var newBarang packets.Barang
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"statusn": 500,
			"error":   err.Error(),
		})
	}

	// extract data
	var data packets.Barang

	if len(form.Value["data"]) == 0 || len(form.File["image"]) == 0 { // check if both data and image is here
		return ctx.Status(500).JSON(fiber.Map{
			"statusn": 500,
			"error":   "Both photo and file object should be in the request",
		})
	}

	err = json.Unmarshal([]byte(form.Value["data"][0]), &data) // decode to data struct object
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"statusn": 500,
			"error":   err.Error(),
		})
	}

	// extract uploaded file
	image := form.File["image"][0]
	imageContent, err := image.Open()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"statusn": 500,
			"error":   err.Error(),
		})
	}

	//	imgBBuffer := bytes.NewBuffer(nil) // create a buffer for the image
	objectName := helper.GenerateToken(image.Filename) + filepath.Ext(image.Filename)

	// preparing a client
	agent := fiber.Put(helper.BarangServiceHostname + "/barang/" + ctx.Params("id"))

	data.Gambar = helper.BarangStoragePublicURL + "/" + objectName // adding gambar URL before sending it
	agent.JSON(data)                                               // push push push

	s, b, e := agent.Bytes()

	if len(e) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  e,
		})
	}

	if s == 200 { // only upload file to the S3 if the data to the database is already inserted

		S3Client := helper.S3Connect()
		uploadCTX := context.Background() //S3Context

		_, err = S3Client.PutObject(uploadCTX, helper.BarangBucketName, objectName, imageContent, image.Size, minio.PutObjectOptions{
			ContentType: image.Header.Get("Content-Type"),
		})
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"statusn": 500,
				"error":   "failed to upload data, an error occured : " + err.Error(),
			})
		}
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(s).Send(b)
}

func ReqDeleteBarang(ctx *fiber.Ctx) error {
	agent := fiber.Delete(helper.BarangServiceHostname + "/barang/" + ctx.Params("id"))
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
