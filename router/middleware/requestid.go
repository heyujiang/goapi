package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func Requestid() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		requestId := ctx.Request.Header.Get("X-Request-Id")

		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}

		ctx.Set("X-Request-Id", requestId)
		ctx.Writer.Header().Set("X-Request-Id", requestId)
		ctx.Next()
	}
}
