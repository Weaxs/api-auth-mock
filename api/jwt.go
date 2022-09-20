package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"strings"
)

const (
	tokenHeaderPrefix = "Bearer "
)

var (
	hmacKeyPath      = "conf/hmac_key"
	rsaPublicKeyPath = "conf/public_key.pub"
)

func RsaApi(c context.Context, ctx *app.RequestContext) {
	token := ctx.Request.Header.Get(authHeader)
	if token == "" || !strings.HasPrefix(token, tokenHeaderPrefix) {
		unauthorized(ctx, "authorization not found.")
		return
	}
	token = strings.TrimPrefix(token, tokenHeaderPrefix)
	pubFile, _ := ioutil.ReadFile(rsaPublicKeyPath)
	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(pubFile)

	claims, err := jwtParse(c, ctx, token, pubKey)
	if err != nil {
		unauthorized(ctx, err.Error())
		return
	}
	key := claims["key"]
	if key != "api-mock-rsa" {
		unauthorized(ctx, "validate authorization failed.")
		return
	}
	success(ctx)
}

func HmacApi(c context.Context, ctx *app.RequestContext) {
	token := ctx.Request.Header.Get(authHeader)
	if token == "" || !strings.HasPrefix(token, tokenHeaderPrefix) {
		unauthorized(ctx, "authorization not found.")
		return
	}
	token = strings.TrimPrefix(token, tokenHeaderPrefix)

	hmacKey, _ := ioutil.ReadFile(hmacKeyPath)
	claims, err := jwtParse(c, ctx, token, hmacKey)
	if err != nil {
		return
	}
	key := claims["key"]
	if key != "api-mock-hmac" {
		unauthorized(ctx, "validate authorization failed.")
		return
	}
	success(ctx)
}

func jwtParse(c context.Context, ctx *app.RequestContext, token string, key interface{}) (jwt.MapClaims, error) {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		unauthorized(ctx, err.Error())
		return nil, err
	}
	return parse.Claims.(jwt.MapClaims), nil
}
