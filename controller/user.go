package controller

import (
	"github.com/gin-gonic/gin"
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
		SendError(ctx, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	//验证数据
	if err = u.Validate(); err != nil {
		SendError(ctx, errno.ErrVaildation, nil)
		return
	}

	//用户密码加密
	if err = u.Encrypt(); err != nil {
		SendError(ctx, errno.ErrEncrypt, nil)
		return
	}

	//创建用户
	if err := u.Create(); err != nil {
		SendError(ctx, errno.ErrCreateUser, nil)
		return
	}

	rsp := UserResponse{
		Username: r.Username,
	}

	SendSuccess(ctx, rsp)
	return
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
		SendError(ctx, errno.ErrDeleteUser, nil)
		return
	}

	SendSuccess(ctx, nil)
	return
}

//根据Id获得用户信息
func GetUserInfo(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))

	user, err := model.GetUser(uint64(userId))

	if err != nil {
		SendError(ctx, errno.ErrData, nil)
		return
	}

	SendSuccess(ctx, user)
	return
}

//更新用户
func UpdateUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))

	var userModel model.UserModel

	if err := ctx.Bind(&userModel); err != nil {
		SendError(ctx, err, nil)
	}

	userModel.Id = uint64(userId)

	if err := userModel.Validate(); err != nil {
		SendError(ctx, err, nil)
	}

	if err := userModel.Encrypt(); err != nil {
		SendError(ctx, err, nil)
	}

	if err := userModel.Update(); err != nil {
		SendError(ctx, err, nil)
	}

	SendSuccess(ctx, nil)

}

//用户列表
func UserList(ctx *gin.Context) {

	var listRequest ListRequest

	if err := ctx.ShouldBind(&listRequest); err != nil {
		SendError(ctx, err, nil)
	}

	list, count, err := service.ListUser(listRequest.Offset, listRequest.Limit)
	if err != nil {
		SendError(ctx, err, nil)
		return
	}

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	SendSuccess(ctx, res)
	return

}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var loginRequest LoginRequest

	if err := ctx.ShouldBind(&loginRequest); err != nil {
		SendError(ctx, err, nil)
		return
	}

	tokenString, err := login.Login(ctx, loginRequest.Username, loginRequest.Password)
	if err != nil {
		SendError(ctx, err, nil)
		return
	}

	SendSuccess(ctx, model.Token{Token: tokenString})
	return
}
