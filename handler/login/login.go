package login

import (
	"github.com/gin-gonic/gin"
	"goapi/handler"
	"goapi/model"
	"goapi/pkg/errno"
	"goapi/service/login"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var loginRequest LoginRequest

	if err := ctx.ShouldBind(&loginRequest); err != nil {
		handler.SendResponse(ctx, err, nil)
		return
	}

	tokenString, err := login.Login(ctx, loginRequest.Username, loginRequest.Password)
	if err != nil {
		handler.SendResponse(ctx, err, nil)
		return
	}

	handler.SendResponse(ctx, errno.LoginSuccess, model.Token{Token: tokenString})
}

func Logout(ctx *gin.Context) {

}
