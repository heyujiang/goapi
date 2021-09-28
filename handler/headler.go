package handler

import (
	"github.com/gin-gonic/gin"
	"goapi/pkg/errno"
	"net/http"
)

type Response struct {
	Code int
	Message string
	Data interface{}
}


func SendResponse(ctx *gin.Context,err error,data interface{}) {
	code , message := errno.DecodeErr(err)
	ctx.JSON(http.StatusOK,Response{
		Code:code,
		Message:message,
		Data:data,
	})
}