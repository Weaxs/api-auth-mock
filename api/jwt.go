package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strings"
)

const (
	tokenHeaderPrefix = "Bearer "
)

var (
	hmacKeyPath        = "conf/hmac_key"
	rsaPublicKeyPath   = "conf/public_key.pub"
	es256PublicKeyPath = "conf/ec256-public.pem"
	es384PublicKeyPath = "conf/ec384-public.pem"
	es512PublicKeyPath = "conf/ec512-public.pem"
)

type algError struct {
	alg string
}

func (algError *algError) Error() string {
	return algError.alg + " is not support."
}

func RsaApi(c context.Context, ctx *app.RequestContext) {
	token := ctx.Request.Header.Get(authHeader)
	if token == "" || !strings.HasPrefix(token, tokenHeaderPrefix) {
		unauthorized(ctx, "authorization not found.")
		return
	}
	token = strings.TrimPrefix(token, tokenHeaderPrefix)

	claims, err := jwtParse(c, ctx, token, func(token *jwt.Token) (interface{}, error) {
		alg := token.Header["alg"].(string)
		if alg != "RS256" && alg != "RS384" && alg != "RS512" {
			return nil, &algError{alg: alg}
		}
		pubFile, _ := os.ReadFile(rsaPublicKeyPath)
		pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(pubFile)
		return pubKey, nil
	})
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

	claims, err := jwtParse(c, ctx, token, func(token *jwt.Token) (interface{}, error) {
		alg := token.Header["alg"].(string)
		if alg != "HS256" && alg != "HS384" && alg != "HS512" {
			return nil, &algError{alg: alg}
		}
		hmacKey, _ := os.ReadFile(hmacKeyPath)
		return hmacKey, nil
	})
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

func EcdsaApi(c context.Context, ctx *app.RequestContext) {
	token := ctx.Request.Header.Get(authHeader)
	if token == "" || !strings.HasPrefix(token, tokenHeaderPrefix) {
		unauthorized(ctx, "authorization not found.")
		return
	}
	token = strings.TrimPrefix(token, tokenHeaderPrefix)

	claims, err := jwtParse(c, ctx, token, func(token *jwt.Token) (interface{}, error) {
		alg := token.Header["alg"].(string)
		var key []byte
		if alg == "ES256" {
			key, _ = os.ReadFile(es256PublicKeyPath)
		} else if alg == "ES384" {
			key, _ = os.ReadFile(es384PublicKeyPath)
		} else if alg == "ES512" {
			key, _ = os.ReadFile(es512PublicKeyPath)
		} else {
			return nil, &algError{alg: alg}
		}
		pub, err := jwt.ParseECPublicKeyFromPEM(key)
		if err != nil {
			return nil, err
		}
		return pub, nil
	})
	if err != nil {
		return
	}
	key := claims["key"]
	if key != "api-mock-ecdsa" {
		unauthorized(ctx, "validate authorization failed.")
		return
	}
	success(ctx)
}

func jwtParse(c context.Context, ctx *app.RequestContext, token string, keyfunc jwt.Keyfunc) (jwt.MapClaims, error) {
	parse, err := jwt.Parse(token, keyfunc)
	if err != nil {
		unauthorized(ctx, err.Error())
		return nil, err
	}
	return parse.Claims.(jwt.MapClaims), nil
}
