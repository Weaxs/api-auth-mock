package api

import (
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/route"
	"strings"
	"testing"
)

func TestOauth(t *testing.T) {
	router := route.NewEngine(config.NewOptions([]config.Option{}))
	router.GET("/api/oauth/mock/token", OauthToken)
	router.GET("/api/oauth/mock/authorize", OauthAuth)
	RegisterOauthServer()

	var token string
	t.Run("token", func(t *testing.T) {
		params := "?grant_type=client_credentials&client_id=id0001&client_secret=secret0001"
		w := ut.PerformRequest(router, "GET", "/api/oauth/mock/token"+params, nil, ut.Header{})
		resp := w.Result()
		var body map[string]any
		err := json.Unmarshal(resp.Body(), &body)
		if err != nil {
			t.Log(err.Error())
			t.Fail()
		}
		assert.NotNil(t, body)
		assert.True(t, "Bearer" == body["token_type"].(string))
		token = body["token_type"].(string) + " " + body["access_token"].(string)
	})

	t.Run("authorize", func(t *testing.T) {
		params := "?response_type=code&client_id=id0001&redirect_uri=https://google.com"
		w := ut.PerformRequest(router, "GET", "/api/oauth/mock/authorize"+params, nil,
			ut.Header{Key: authHeader, Value: token})
		resp := w.Result()
		assert.DeepEqual(t, 302, resp.StatusCode())
		assert.True(t, strings.HasPrefix(string(resp.Header.PeekLocation()), "https://google.com"))
	})

}
