package system

import (
	"github.com/ducconit/gobase/api/context"
	"net/http"
)

func Ping(ctx context.Context) error {
	return ctx.JSON(http.StatusOK, map[string]any{
		"status":  200,
		"message": "pong",
	})
}
