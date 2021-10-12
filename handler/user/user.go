package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/handler"
	"goapi/model"
	"goapi/pkg/errno"
	"strconv"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Username string `json:"username"`
}

//创建新用户
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

//删除用户
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

//根据Id获得用户信息
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

//更新用户
func Update(ctx *gin.Context) {

}

//用户列表
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
