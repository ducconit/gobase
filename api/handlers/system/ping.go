package system

import (
	"github.com/gofiber/fiber/v2"
)

func Ping(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status":  200,
		"message": "pong",
	})
}
