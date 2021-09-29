package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/handler"
	"goapi/model"
	"goapi/pkg/errno"
	"strconv"
)

func Get(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))

	user, err := model.GetUser(uint64(userId))

	if err != nil {
		fmt.Println(err.Error())
		handler.SendResponse(ctx, errno.ErrData, nil)
		return
	}

	handler.SendResponse(ctx, nil, user)
}
