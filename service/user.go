package service

import (
	"fmt"
	"goapi/model"
	"goapi/util"
	"sync"
)

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
