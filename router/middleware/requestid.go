package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Requestid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("ctx *gin.Context")
		ctx.Next()
	}
}
