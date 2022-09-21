package main

import (
	"github.com/Weaxs/api-auth-mock/api"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterGroupRoute(h *server.Hertz) {
	apiGroup := h.Group("/api")

	// jwt api
	jwtGroup := apiGroup.Group("/jwt/mock")
	{
		jwtGroup.GET("/hmac", api.HmacApi)
		jwtGroup.GET("/rsa", api.RsaApi)
		jwtGroup.GET("/ecdsa", api.EcdsaApi)
	}

	// basic api
	basic := apiGroup.Group("/basic/mock")
	{
		// basic middleware
		basic.Use(basic_auth.BasicAuth(map[string]string{
			"account1": "password1",
			"account2": "password2",
			"account3": "password3",
		}))
		basic.GET("", api.BasicAuthApi)
	}

	// oauth api
	api.RegisterOauthServer()
	oauth := apiGroup.Group("/oauth/mock")
	{
		oauth.GET("/token", api.OauthToken)
		oauth.GET("/authorize", api.OauthAuth)
	}
}
