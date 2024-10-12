package context

import "github.com/gofiber/fiber/v2"

type Context struct {
	*fiber.Ctx
}

func NewContext(ctx *fiber.Ctx) error {
	c := &Context{ctx}

	c.Locals("ctx", c)

	return ctx.Next()
}
