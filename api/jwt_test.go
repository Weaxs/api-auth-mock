package api

import (
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/dgrijalva/jwt-go"
	"os"
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

func TestEcdsa(t *testing.T) {
	t.Run("ecdsa256", func(t *testing.T) {
		es256PublicKeyPath = "../conf/ec256-public.pem"
		ecdsa(t, jwt.SigningMethodES256, "../conf/ec256-private.pem")
	})
	t.Run("ecdsa384", func(t *testing.T) {
		es384PublicKeyPath = "../conf/ec384-public.pem"
		ecdsa(t, jwt.SigningMethodES384, "../conf/ec384-private.pem")
	})
	t.Run("ecdsa512", func(t *testing.T) {
		es512PublicKeyPath = "../conf/ec512-public.pem"
		ecdsa(t, jwt.SigningMethodES512, "../conf/ec512-private.pem")
	})
}

func hmac(t *testing.T, method jwt.SigningMethod) {
	hmacKeyPath = "../conf/hmac_key"
	router := route.NewEngine(config.NewOptions([]config.Option{}))
	router.GET("/api/jwt/mock/hmac", HmacApi)

	claims := jwt.MapClaims{
		"key": "api-mock-hmac",
	}
	hmacKey, _ := os.ReadFile("../conf/hmac_key")
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
	privateFile, _ := os.ReadFile("../conf/private_key")
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

func ecdsa(t *testing.T, method jwt.SigningMethod, privateKeyPath string) {
	path := "/api/jwt/mock/ecdsa"
	router := route.NewEngine(config.NewOptions([]config.Option{}))
	router.GET(path, EcdsaApi)

	claims := jwt.MapClaims{
		"key": "api-mock-ecdsa",
	}
	privateFile, _ := os.ReadFile(privateKeyPath)
	privateKey, _ := jwt.ParseECPrivateKeyFromPEM(privateFile)
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
