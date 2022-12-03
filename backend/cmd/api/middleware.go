package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := ctx.Cookie("auth")
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
