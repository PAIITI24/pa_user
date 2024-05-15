package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/pa_user/model"
	"time"
)

/**
We will do token checking over here.
*/

func token_login(ctx *fiber.Ctx) error {
	authorizationKey := ctx.Accepts("Authorization") // get the used auth key

	// try to find the key data
	var tokendata model.Token
	fetched := db.Where(&model.Token{Token: authorizationKey}).First(&tokendata).RowsAffected

	// if none
	if fetched == 0 && tokendata.ExpiredAt.Before(time.Now()) && tokendata.IsEnabled {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": fiber.StatusUnauthorized,
			"error":  "Your token isn't valid or already expired",
		})
	}

	return ctx.Next()
}
