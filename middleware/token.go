package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hakushigo/pa_user/model"
	"time"
)

/**
We will do token checking over here.
*/

func TokenLogin(ctx *fiber.Ctx) error {
	authorizationKey := ctx.Get("Authorization") // get the used auth key
	if authorizationKey == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": fiber.StatusUnauthorized,
			"error":  "Authorization is needed",
		})
	}

	// try to find the key data
	var tokendata model.Token
	find := db.Where(&model.Token{Token: authorizationKey}).Find(&tokendata).RowsAffected

	// if none
	if find == 0 || tokendata.ExpiredAt.Before(time.Now()) || !tokendata.IsEnabled {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": fiber.StatusUnauthorized,
			"error":  "Your token isn't valid or already expired",
		})
	} else {
		return ctx.Next()
	}
}
