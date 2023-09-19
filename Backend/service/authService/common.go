package authservice

import (
	"sync"
)

var UserServiceIns *UserService
var UserServiceOnce sync.Once

type UserService struct {
}

func GetUserService() *UserService {
	UserServiceOnce.Do(func() {
		UserServiceIns = &UserService{}
	})
	return UserServiceIns
}
