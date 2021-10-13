package middleware

import (
	"github.com/gin-gonic/gin"
	"goapi/handler"
	"goapi/pkg/errno"
	"goapi/pkg/token"
)

func AuthMiddleware(ctx *gin.Context) {
	//验证token
	if _, err := token.ParseRequest(ctx); err != nil {
		handler.SendResponse(ctx, errno.ErrTokenInvalid, nil)
		ctx.Abort()
		return
	}
	ctx.Next()
}
