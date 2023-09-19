package authservice

import (
	"context"
	"errors"
	"resume_backend/models"
	"resume_backend/pkg/ctl"
	"resume_backend/pkg/utils"
	"resume_backend/repositories"

	"gorm.io/gorm"
)

func (s *UserService) UserRegister(ctx context.Context, req *models.UserServiceRegisterReq) (resp interface{}, err error) {
	userDao := repositories.NewUserDao(ctx)
	u, err := userDao.FindUserByUserName(req.UserName)
	switch err {
	case gorm.ErrRecordNotFound:
		u = &models.User{
			UserName: req.UserName,
		}
		// 密码加密存储
		if err = u.SetPassword(req.Password); err != nil {
			utils.LogrusObj.Info(err)
			return
		}
		if err = userDao.CreateUser(u); err != nil {
			utils.LogrusObj.Info(err)
			return
		}
		return ctl.RespSuccess(), nil
	case nil:
		err = errors.New("用户已存在")
		return
	default:
		return
	}
}
