package login

import (
	"goapi/model"
	"goapi/pkg/errno"
)

func Login(username, password string) error {
	//根据用户名获得用户信息
	userModel, err := model.GetUserByUserName(username)
	if err != nil {
		return err
	}

	if userModel == nil {
		return errno.NoUsername
	}

	if err := userModel.Compare(password); err != nil {
		return err
	}

	return errno.LoginSuccess
}
