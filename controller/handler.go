package controller

import (
	"github.com/gin-gonic/gin"
	"goapi/pkg/errno"
	"net/http"
)

//接口返回数据
type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(ctx *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	ctx.JSON(http.StatusOK, response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

//操作成功，返回数据
func SendSuccess(ctx *gin.Context, data interface{}) {
	SendResponse(ctx, errno.SUCCESS, data)
}

//操作失败
func SendError(ctx *gin.Context, err error, data interface{}) {
	SendResponse(ctx, err, data)
}
