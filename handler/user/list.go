package user

import (
	"github.com/gin-gonic/gin"
	"goapi/handler"
	"goapi/model"
)

func List(ctx *gin.Context) {
	list, count, err := model.ListUser(1, 30)
	if err != nil {
		handler.SendResponse(ctx, err, nil)
		return
	}

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	handler.SendResponse(ctx, nil, res)
	return

}
