package user

import (
	"github.com/gin-gonic/gin"
	"goapi/handler"
	"goapi/model"
	"goapi/pkg/errno"
)

func Create(ctx *gin.Context) {
	var r UserRequest

	var err error
	if err := ctx.Bind(&r); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	//验证数据
	if err = u.Validate(); err != nil {
		handler.SendResponse(ctx, errno.ErrVaildation, nil)
		return
	}

	//用户密码加密
	if err = u.Encrypt(); err != nil {
		handler.SendResponse(ctx, errno.ErrEncrypt, nil)
		return
	}

	//创建用户
	if err := u.Create(); err != nil {
		handler.SendResponse(ctx, errno.ErrCreateUser, nil)
		return
	}

	rsp := UserResponse{
		Username: r.Username,
	}

	handler.SendResponse(ctx, nil, rsp)

}
