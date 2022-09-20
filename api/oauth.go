package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/go-oauth2/oauth2/manage"
	"github.com/go-oauth2/oauth2/models"
	"github.com/go-oauth2/oauth2/server"
	"github.com/go-oauth2/oauth2/store"
)

var oauthServ *server.Server

// RegisterOauthServer init oauth server
func RegisterOauthServer() {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()
	_ = clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
	})
	manager.MapClientStorage(clientStore)

	oauthServ = server.NewDefaultServer(manager)
	oauthServ.SetAllowGetAccessRequest(true)
	oauthServ.SetClientInfoHandler(server.ClientFormHandler)
}

// OauthToken get oauth token
func OauthToken(c context.Context, ctx *app.RequestContext) {
	request, _ := adaptor.GetCompatRequest(&ctx.Request)
	responseWriter := adaptor.GetCompatResponseWriter(&ctx.Response)

	_ = oauthServ.HandleTokenRequest(responseWriter, request)
}

// OauthAuth authorize token
func OauthAuth(c context.Context, ctx *app.RequestContext) {
	request, _ := adaptor.GetCompatRequest(&ctx.Request)
	responseWriter := adaptor.GetCompatResponseWriter(&ctx.Response)

	err := oauthServ.HandleAuthorizeRequest(responseWriter, request)
	if err != nil {
		unauthorized(ctx, err.Error())
		return
	}
}
