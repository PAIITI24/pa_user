package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/pa_user/controller"
	service "github.com/hakushigo/pa_user/controller/other_services"
	"github.com/hakushigo/pa_user/helper"
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
	user := server.Group("/user")
	user.Post("/signup", controller.CreateUser)
	user.Post("/delete", controller.DeleteUser)

	// kategori obat
	obat := server.Group("/obat")
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

	// kategori produk
	produk := server.Group("/obat")
	produk.Post("/kategori/", nil)
	produk.Get("/kategori/:id", nil)
	produk.Get("/kategori/", nil)
	produk.Put("/kategori/:id", nil)
	produk.Delete("/kategori/:id", nil)
	produk.Get("/kategori/count", nil)

	// obat
	produk.Post("/", service.ReqAddObat)
	produk.Get("/", service.ReqListObat)
	produk.Get("/:id", service.ReqGetObat)
	produk.Put("/:id", service.ReqUpdateObat)
	produk.Delete("/:id", service.ReqDeleteObat)

	server.Listen(":3003")
}
