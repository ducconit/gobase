package system

import (
	"github.com/ducconit/gobase/api/context"
	"github.com/gofiber/fiber/v2"
)

func Ping(ctx *context.Context) error {
	return ctx.JSON(fiber.Map{
		"status":  200,
		"message": "pong",
	})
}
