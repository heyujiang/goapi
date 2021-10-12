package login

import (
	"github.com/gin-gonic/gin"
	"goapi/handler"
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
	}

	handler.SendResponse(ctx, login.Login(loginRequest.Username, loginRequest.Password), nil)
}

func Logout(ctx *gin.Context) {

}
