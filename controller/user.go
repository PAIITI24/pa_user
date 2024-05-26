package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/pa_user/helper"
	"github.com/hakushigo/pa_user/model"
	"time"
)

func DeleteUser(ctx *fiber.Ctx) error {
	Token := ctx.Get("Authorization")

	var userdata model.User

	err := ctx.BodyParser(&userdata)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	// get user data of the one that asked
	var userwhodelete model.User

	err = db.Preload("Token").Where("token = ? AND expired_at", Token, time.Now()).First(&userwhodelete).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	// and then check if the user based on the user id has owner priviledged.
	if userwhodelete.Role == 99 { // if the role is owner (0)
		err = db.Find(&userdata).Delete(&userdata).Error
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"status": 500,
				"error":  err,
			})
		}

		return ctx.Status(200).JSON(fiber.Map{
			"status":       200,
			"deleted_user": userdata,
		})
	} else { // if anything else
		return nil
	}
}

func Login(ctx *fiber.Ctx) error {
	// variable to store login information
	var logininfo model.User

	// parse body
	err := ctx.BodyParser(&logininfo)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	var foundUser model.User
	search := db.Model(&model.User{}).Where(&model.User{Email: logininfo.Email}).Find(&foundUser)

	if search.RowsAffected >= 0 && helper.HashCheck(logininfo.Password, foundUser.Password) {
		// after that, we go to the next step which is add a new token
		var newToken model.Token = model.Token{
			Token:     helper.GenerateToken(foundUser.Password),
			Owner:     foundUser.Id,                        // put user object as user
			ExpiredAt: time.Now().Add(30 * 24 * time.Hour), // 30 days from now
			IsEnabled: true,
		}

		err = db.Create(&newToken).Error // store the data

		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"status": 500,
				"error":  err,
			})
		} else {
			return ctx.Status(200).JSON(fiber.Map{
				"status": 200,
				"token":  newToken.Token,
				"user": fiber.Map{
					"email": foundUser.Email,
					"role":  foundUser.Role,
				},
			})
		}
	} else {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Sorry, please check your email & password",
		})
	}
}

func CreateUser(ctx *fiber.Ctx) error {
	Token := ctx.Get("Authorization")
	var userdata model.User

	err := ctx.BodyParser(&userdata)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	// get user data of the one that asked
	var userwhoadd model.User
	var tokendata model.Token

	find := db.Model(&model.Token{}).Preload("User").Where(model.Token{
		Token: Token,
	}).Find(&tokendata)

	if find.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	userwhoadd = tokendata.User

	// and then check if the user based on the user id has owner priviledged.
	if userwhoadd.Role == 99 { // if the role is owner (0)
		userdata.Role = 1 // override the role settings

		if userdata.Password != "" && userdata.Email != "" {
			userdata.Password, err = helper.PasswordHash(userdata.Password)
			err = db.Create(&userdata).Error

			if err != nil {
				return ctx.Status(500).JSON(fiber.Map{
					"status": 500,
					"error":  err,
				})
			}

			return ctx.Status(200).JSON(fiber.Map{
				"status": 200,
				"user":   userdata,
			})
		} else {
			return ctx.Status(500).JSON(fiber.Map{
				"status": 500,
				"error":  "the user's email or password is empty",
			})
		}

	} else { // if anything else
		return ctx.Status(403).JSON(fiber.Map{
			"status":  403,
			"message": "You are not allowed to add this entry",
		})
	}
}

func Logout(ctx *fiber.Ctx) error {
	var AuthenticationToken = ctx.Get("Authorization")
	var TokenData model.Token

	search := db.Where(&model.Token{Token: AuthenticationToken}).Find(&TokenData)

	if search.RowsAffected == 0 || !TokenData.IsEnabled { // if not found or disabled
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"token":   TokenData.Token,
			"message": "Your token isn't valid",
		})
	} else {

		TokenData.IsEnabled = false
		db.Save(&TokenData)

		return ctx.Status(200).JSON(fiber.Map{
			"status": 200,
			"token": fiber.Map{
				"token":  TokenData.Token,
				"status": TokenData.IsEnabled,
			},
		})
	}
}
