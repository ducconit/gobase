package context

import "github.com/gofiber/fiber/v2"

type Context struct {
	*fiber.Ctx
}

type Handler func(*Context) error

func RegisterContextFiber(ctx *fiber.Ctx) error {
	ctx.Locals("context", &Context{Ctx: ctx})
	return ctx.Next()
}

func WithContextFiber(h Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Locals("context").(*Context)
		return h(ctx)
	}
}
