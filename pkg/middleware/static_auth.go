package middleware

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type StaticAuth struct {
	login    string
	password string
	realm    string
}

func (auth *StaticAuth) Handle(ctx *gin.Context) {
	login, password, ok := ctx.Request.BasicAuth()
	if !ok || auth.login != login || auth.password != password {
		ctx.Status(http.StatusUnauthorized)
		ctx.Header("WWW-Authenticate", `Basic realm="`+auth.realm+`", charset="UTF-8"`)
		ctx.Abort()
	}
}

func NewStaticAuth(login, password, realm string) *StaticAuth {
	return &StaticAuth{
		login:    login,
		password: password,
		realm:    url.QueryEscape(realm),
	}
}
