package login

import (
	"github.com/gin-gonic/gin"
	"goapi/model"
	"goapi/pkg/errno"
	"goapi/pkg/token"
)

func Login(ctx *gin.Context, username, password string) (string, error) {
	//根据用户名获得用户信息
	userModel, err := model.GetUserByUserName(username)
	if err != nil {
		return "", err
	}

	if userModel == nil {
		return "", errno.NoUsername
	}

	if err := userModel.Compare(password); err != nil {
		return "", err
	}

	tokenString, err := token.Sign(ctx, &token.Context{userModel.Id, userModel.Username})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
