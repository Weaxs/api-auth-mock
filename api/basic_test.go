package api

import (
	"encoding/base64"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/route"
	"testing"
)

func TestBasic(t *testing.T) {
	router := route.NewEngine(config.NewOptions([]config.Option{}))
	router.GET("/api/basic/mock", BasicAuthApi)

	token := "Basic " + base64.StdEncoding.EncodeToString([]byte("account1:password1"))
	w := ut.PerformRequest(router, "GET", "/api/basic/mock", nil,
		ut.Header{"Authorization", token})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}
