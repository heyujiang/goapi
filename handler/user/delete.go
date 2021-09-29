package user

import (
	"github.com/gin-gonic/gin"
	"goapi/handler"
	"goapi/model"
	"goapi/pkg/errno"
	"strconv"
)

func Delete(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))

	user := model.UserModel{
		BaseModel: model.BaseModel{
			Id: uint64(userId),
		},
	}

	if err := user.Delete(); err != nil {
		handler.SendResponse(ctx, errno.ErrDeleteUser, nil)
	}

	handler.SendResponse(ctx, nil, nil)
}
