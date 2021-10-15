package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/handler"
	"goapi/model"
	"goapi/pkg/errno"
	"goapi/service"
	"goapi/service/login"
	"strconv"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

//创建新用户
func CreateUser(ctx *gin.Context) {
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
func DeleteUser(ctx *gin.Context) {
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
func GetUserInfo(ctx *gin.Context) {
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
func UpdateUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))

	var userModel model.UserModel

	if err := ctx.Bind(&userModel); err != nil {
		handler.SendResponse(ctx, err, nil)
	}

	userModel.Id = uint64(userId)

	if err := userModel.Validate(); err != nil {
		handler.SendResponse(ctx, err, nil)
	}

	if err := userModel.Encrypt(); err != nil {
		handler.SendResponse(ctx, err, nil)
	}

	if err := userModel.Update(); err != nil {
		handler.SendResponse(ctx, err, nil)
	}

	handler.SendResponse(ctx, errno.OK, nil)

}

//用户列表
func UserList(ctx *gin.Context) {

	var listRequest ListRequest

	if err := ctx.ShouldBind(&listRequest); err != nil {
		handler.SendResponse(ctx, err, nil)
	}

	list, count, err := service.ListUser(listRequest.Offset, listRequest.Limit)
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