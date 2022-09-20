package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func BasicAuthApi(c context.Context, ctx *app.RequestContext) {
	success(ctx)
}
