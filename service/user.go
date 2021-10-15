package service

import (
	"fmt"
	"goapi/entity/bo"
	"goapi/model"
	"goapi/pkg/errno"
	"goapi/pkg/token"
	"goapi/util"
	"sync"
)

func CreateUser(so *bo.CreateUserBo) error {
	u := model.UserModel{
		Username: so.Username,
		Password: so.Password,
	}

	//验证数据
	if err := u.Validate(); err != nil {
		return err
	}

	//用户密码加密
	if err := u.Encrypt(); err != nil {
		return errno.ErrEncrypt
	}

	//创建用户
	if err := u.Create(); err != nil {
		return errno.ErrCreateUser
	}

	return nil
}

func DeleteUserById(userId uint64) error {
	user := model.UserModel{
		BaseModel: model.BaseModel{
			Id: userId,
		},
	}

	if err := user.Delete(); err != nil {
		return errno.ErrDeleteUser
	}

	return nil
}

func GetUserInfoById(userId uint64) (*model.UserModel, error) {
	user, err := model.GetUser(uint64(userId))

	if err != nil {
		return user, err
	}

	return user, nil
}

func ListUser(offset, limit int) ([]model.UserInfo, uint64, error) {

	users, count, err := model.ListUser(offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}
	finished := make(chan bool)
	errChan := make(chan error)

	type UserList struct {
		Lock   *sync.Mutex
		IdMaps map[uint64]model.UserInfo
	}

	userList := UserList{
		Lock:   new(sync.Mutex),
		IdMaps: make(map[uint64]model.UserInfo, len(users)),
	}

	for _, user := range users {
		wg.Add(1)
		go func(user *model.UserModel) {
			defer wg.Done()

			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMaps[user.Id] = model.UserInfo{
				Id:        user.Id,
				Username:  user.Username,
				Password:  user.Password,
				SayHello:  fmt.Sprintf("Hello %s", shortId),
				CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
			}

		}(user)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, 0, err
	}

	infos := make([]model.UserInfo, 0)
	for _, id := range ids {
		infos = append(infos, userList.IdMaps[id])
	}

	return infos, count, nil
}

//用户登录
func Login(username, password string) (string, error) {
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

	tokenString, err := token.Sign(&token.Context{userModel.Id, userModel.Username})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
