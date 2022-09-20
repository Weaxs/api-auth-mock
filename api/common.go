package api

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

const authHeader = "Authorization"

func unauthorized(ctx *app.RequestContext, message string) {
	ctx.JSON(consts.StatusUnauthorized, map[string]any{
		"status":  false,
		"message": "Unauthorized: " + message,
	})
}

func success(ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, map[string]any{
		"status":  true,
		"message": "Success.",
	})
}
