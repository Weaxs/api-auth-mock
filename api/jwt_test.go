package api

import (
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"testing"
)

func TestHmac(t *testing.T) {
	t.Run("hmac256", func(t *testing.T) {
		hmac(t, jwt.SigningMethodHS256)
	})
	t.Run("hmac384", func(t *testing.T) {
		hmac(t, jwt.SigningMethodHS384)
	})
	t.Run("hmac512", func(t *testing.T) {
		hmac(t, jwt.SigningMethodHS512)
	})
}

func TestRsa(t *testing.T) {
	t.Run("rsa256", func(t *testing.T) {
		rsa(t, jwt.SigningMethodRS256)
	})
	t.Run("rsa384", func(t *testing.T) {
		rsa(t, jwt.SigningMethodRS384)
	})
	t.Run("rsa512", func(t *testing.T) {
		rsa(t, jwt.SigningMethodRS512)
	})
}

func hmac(t *testing.T, method jwt.SigningMethod) {
	hmacKeyPath = "../conf/hmac_key"
	router := route.NewEngine(config.NewOptions([]config.Option{}))
	router.GET("/api/jwt/mock/hmac", HmacApi)

	claims := jwt.MapClaims{
		"key": "api-mock-hmac",
	}
	hmacKey, _ := ioutil.ReadFile("../conf/hmac_key")
	token, err := jwt.NewWithClaims(method, claims).SignedString(hmacKey)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	w := ut.PerformRequest(router, "GET", "/api/jwt/mock/hmac", nil,
		ut.Header{authHeader, tokenHeaderPrefix + token})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func rsa(t *testing.T, method jwt.SigningMethod) {
	rsaPublicKeyPath = "../conf/public_key.pub"
	path := "/api/jwt/mock/rsa"
	router := route.NewEngine(config.NewOptions([]config.Option{}))
	router.GET(path, RsaApi)

	claims := jwt.MapClaims{
		"key": "api-mock-rsa",
	}
	privateFile, _ := ioutil.ReadFile("../conf/private_key")
	privateKey, _ := jwt.ParseRSAPrivateKeyFromPEM(privateFile)
	token, err := jwt.NewWithClaims(method, claims).SignedString(privateKey)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	w := ut.PerformRequest(router, "GET", path, nil,
		ut.Header{authHeader, tokenHeaderPrefix + token})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}
