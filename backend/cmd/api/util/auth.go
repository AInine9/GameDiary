package util

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/objx"
	"net/http"
)

func GetUserFromCookie(ctx *gin.Context) (userId string, userName string) {
	cookie, err := ctx.Cookie("auth")
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		ctx.Abort()
		return
	}
	data := objx.MustFromBase64(cookie)
	userId = data.Get("user_id").String()
	userName = data.Get("user_name").String()
	if userName == "" || userId == "" {
		ctx.Status(http.StatusUnauthorized)
		ctx.Abort()
		return
	}
	return userId, userName
}
