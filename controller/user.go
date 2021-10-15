package controller

import (
	"github.com/gin-gonic/gin"
	"goapi/entity/bo"
	"goapi/entity/dto"
	"goapi/entity/vo"
	"goapi/model"
	"goapi/pkg/errno"
	"goapi/service"
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
	var createUserDto dto.CreateUserDto

	if err := ctx.Bind(&createUserDto); err != nil {
		SendError(ctx, errno.ErrBind, nil)
		return
	}

	createUserBo := &bo.CreateUserBo{
		Username: createUserDto.Username,
		Password: createUserDto.Password,
	}

	if err := service.CreateUser(createUserBo); err != nil {
		SendError(ctx, err, nil)
		return
	}

	SendSuccess(ctx, vo.CreateUserVo{createUserDto.Username})
	return
}

//删除用户
func DeleteUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))

	if err := service.DeleteUserById(uint64(userId)); err != nil {
		SendError(ctx, err, nil)
		return
	}

	SendSuccess(ctx, nil)
	return
}

//根据Id获得用户信息
func GetUserInfo(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))

	user, err := service.GetUserInfoById(uint64(userId))
	if err != nil {
		SendError(ctx, err, nil)
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

//用户登录
func Login(ctx *gin.Context) {
	var loginDto dto.LoginDto

	if err := ctx.ShouldBind(&loginDto); err != nil {
		SendError(ctx, err, nil)
		return
	}

	tokenString, err := service.Login(loginDto.Username, loginDto.Password)
	if err != nil {
		SendError(ctx, err, nil)
		return
	}

	SendSuccess(ctx, model.Token{Token: tokenString})
	return
}
