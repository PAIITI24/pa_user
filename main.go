package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/pa_user/controller"
	service "github.com/hakushigo/pa_user/controller/other_services"
	"github.com/hakushigo/pa_user/helper"
	"github.com/hakushigo/pa_user/middleware"
	"github.com/hakushigo/pa_user/model"
	"time"
)

func main() {

	// Migrate table
	helper.Migrator()

	db := helper.DB()

	// create owner at launch
	// search if owner already exists
	if db.Find(&model.User{}).Where("role", 0).RowsAffected == 0 {

		password, _ := helper.PasswordHash("abc123")

		user := model.User{
			Role:     99,
			Email:    "admin@admin.com",
			Password: password,
			Name:     "John",
		}

		token := model.Token{
			Token:     "abc123",
			IsEnabled: true,
			ExpiredAt: time.Now().Add(30 * 24 * time.Hour),
			User:      user,
		}

		db.Create(&token)
	}

	// declare server
	server := fiber.New(
		fiber.Config{
			Immutable: false,
			AppName:   "User_Apotek_APP",
		})

	// auth
	auth := server.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/logout", controller.Logout)

	// user
	user := server.Group("/user", middleware.TokenLogin)
	user.Post("/signup", controller.CreateUser)
	user.Post("/delete", controller.DeleteUser)

	// kategori obat
	obat := server.Group("/obat", middleware.TokenLogin)
	obat.Post("/kategori/", service.ReqAddKategoriObat)
	obat.Get("/kategori/:id", service.ReqGetKategoriObat)
	obat.Get("/kategori/", service.ReqListKategoriObat)
	obat.Put("/kategori/:id", service.ReqUpdateKategoriObat)
	obat.Delete("/kategori/:id", service.ReqDeleteKategoriObat)
	obat.Get("/kategori/count", service.ReqCountKategoriObat)

	// obat
	obat.Post("/", service.ReqAddObat)
	obat.Get("/", service.ReqListObat)
	obat.Get("/:id", service.ReqGetObat)
	obat.Put("/:id", service.ReqUpdateObat)
	obat.Delete("/:id", service.ReqDeleteObat)

	// stok barang
	obat.Put("/stok/add", service.AddStokObat)
	obat.Put("/stok/reduce", service.ReduceStokObat)

	// kategori barang
	barang := server.Group("/barang", middleware.TokenLogin)
	barang.Post("/kategori/", service.ReqAddKategoriBarang)
	barang.Get("/kategori/:id", service.ReqGetKategoriBarang)
	barang.Get("/kategori/", service.ReqListKategoriBarang)
	barang.Put("/kategori/:id", service.ReqUpdateKategoriBarang)
	barang.Delete("/kategori/:id", service.ReqDeleteKategoriBarang)
	barang.Get("/kategori/count", service.ReqCountKategoriBarang)

	// barang
	barang.Post("/", service.ReqAddBarang)
	barang.Get("/", service.ReqListBarang)
	barang.Get("/:id", service.ReqGetBarang)
	barang.Put("/:id", service.ReqUpdateBarang)
	barang.Delete("/:id", service.ReqDeleteBarang)

	// stok barang
	barang.Put("/stok/add", service.AddStokBarang)
	barang.Put("/stok/reduce", service.ReduceStokBarang)

	server.Listen(":3003")
}
