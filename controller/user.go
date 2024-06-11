package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/pa_user/helper"
	"github.com/hakushigo/pa_user/model"
	"time"
)

func DeleteUser(ctx *fiber.Ctx) error {
	Token := ctx.Get("Authorization")
	var tokendata model.Token

	// Get the user ID from the URL parameters
	userID, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid user ID",
		})
	}

	// Get user data of the one that asked
	find := db.Model(&model.Token{}).Preload("User").Where(model.Token{
		Token: Token,
	}).Find(&tokendata)

	if find.RowsAffected == 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "User by this token is not found",
		})
	}

	var userwhodelete model.User = tokendata.User

	// Check if the user has owner privileges
	if userwhodelete.Role == 99 { // if the role is owner
		var userdata model.User
		err = db.Where("id = ?", userID).First(&userdata).Error
		if err != nil {
			return ctx.Status(404).JSON(fiber.Map{
				"status":  404,
				"message": "User not found",
			})
		}

		err = db.Delete(&userdata).Error
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"status": 500,
				"error":  err.Error(),
			})
		}

		return ctx.Status(200).JSON(fiber.Map{
			"status":       200,
			"deleted_user": userdata,
		})
	} else { // if anything else
		return ctx.Status(403).JSON(fiber.Map{
			"status":  403,
			"message": "You are not allowed to delete this user",
		})
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
					"id":    foundUser.Id,
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
					"error":  "Database Error : " + err.Error(),
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

func ListStaff(ctx *fiber.Ctx) error {
	var fetchedUsers []model.User
	var tokendata model.Token

	Token := ctx.Get("Authorization")
	find := db.Model(&model.Token{}).Preload("User").Where(model.Token{
		Token: Token,
	}).Find(&tokendata)

	if find.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  find.Error.Error(),
		})
	}

	if tokendata.User.Role == 99 {
		search := db.Where(model.User{Role: 1}).Find(&fetchedUsers)
		// list the staffs here.
		if search.RowsAffected > 0 || search.Error != nil {
			ctx.Status(500).JSON(fiber.Map{})
		}

		return ctx.Status(200).JSON(fetchedUsers)
	} else {
		return ctx.Status(403).JSON(fiber.Map{
			"status":  403,
			"message": "You are not allowed to access this entry",
		})
	}
}

func SelfUpdateAkun(ctx *fiber.Ctx) error {
	var tokendata model.Token
	Token := ctx.Get("Authorization")

	// Retrieve the updated staff data from the request body
	var updatedUserData model.User
	if err := ctx.BodyParser(&updatedUserData); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	// Fetch the token data and associated user
	find := db.Model(&model.Token{}).Preload("User").Where(model.Token{
		Token: Token,
	}).Find(&tokendata)

	if find.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  find.Error.Error(),
		})
	}

	// Fetch the existing user data from the database
	var existingUser model.User
	search := db.Where(model.User{Id: tokendata.User.Id}).First(&existingUser)
	if search.Error != nil || search.RowsAffected == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  404,
			"message": "User not found",
		})
	}

	// Update the user data
	existingUser.Name = updatedUserData.Name
	existingUser.Email = updatedUserData.Email
	if updatedUserData.Password != "" {
		hashedPassword, err := helper.PasswordHash(updatedUserData.Password)
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"status": 500,
				"error":  "Error hashing password: " + err.Error(),
			})
		}
		existingUser.Password = hashedPassword
	}

	// Save the updated user data to the database
	if err := db.Save(&existingUser).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  "Database error: " + err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": 200,
		"user":   existingUser,
	})
}

func UpdateStaff(ctx *fiber.Ctx) error {
	var tokendata model.Token
	Token := ctx.Get("Authorization")

	// Retrieve the ID of the staff to be updated from the request parameters
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid staff ID",
		})
	}

	// Retrieve the updated staff data from the request body
	var updatedUserData model.User
	if err := ctx.BodyParser(&updatedUserData); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  err,
		})
	}

	// Fetch the token data and associated user
	find := db.Model(&model.Token{}).Preload("User").Where(model.Token{
		Token: Token,
	}).Find(&tokendata)

	if find.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  find.Error.Error(),
		})
	}

	// Check if the user has owner privileges
	if tokendata.User.Role != 99 {
		return ctx.Status(403).JSON(fiber.Map{
			"status":  403,
			"message": "You are not allowed to update this entry",
		})
	}

	// Fetch the existing user data from the database
	var existingUser model.User
	search := db.Where(model.User{Role: 1, Id: id}).First(&existingUser)
	if search.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  404,
			"message": "Staff not found",
		})
	}

	// Update the user data
	existingUser.Name = updatedUserData.Name
	existingUser.Email = updatedUserData.Email
	if updatedUserData.Password != "" {
		hashedPassword, err := helper.PasswordHash(updatedUserData.Password)
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"status": 500,
				"error":  "Error hashing password: " + err.Error(),
			})
		}
		existingUser.Password = hashedPassword
	}

	// Save the updated user data to the database
	if err := db.Save(&existingUser).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  "Database error: " + err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": 200,
		"user":   existingUser,
	})
}

func GetStaff(ctx *fiber.Ctx) error {
	var fetchedUsers model.User
	var tokendata model.Token
	var id, _ = ctx.ParamsInt("id")

	Token := ctx.Get("Authorization")
	find := db.Model(&model.Token{}).Preload("User").Where(model.Token{
		Token: Token,
	}).Find(&tokendata)

	if find.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  find.Error.Error(),
		})
	}

	if tokendata.User.Role == 99 {
		search := db.Where(model.User{Role: 1, Id: id}).First(&fetchedUsers)
		// list the staffs here.
		if search.RowsAffected > 0 || search.Error != nil {
			ctx.Status(500).JSON(fiber.Map{})
		}

		return ctx.Status(200).JSON(fetchedUsers)
	} else {
		return ctx.Status(403).JSON(fiber.Map{
			"status":  403,
			"message": "You are not allowed to access this entry",
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
