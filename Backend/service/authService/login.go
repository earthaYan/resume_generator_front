package authservice

import (
	"context"
	"errors"
	"resume_backend/models"
	"resume_backend/pkg/ctl"
	"resume_backend/pkg/utils"
	"resume_backend/repositories"
)

func (s *UserService) UserLogin(ctx context.Context, req *models.UserServiceLoginReq) (resp interface{}, err error) {
	userDao := repositories.NewUserDao(ctx)
	u, err := userDao.FindUserByUserName(req.UserName)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	// 校验密码
	if !u.CheckPwd(req.Password) {
		err = errors.New("账号或密码错误")
		utils.LogrusObj.Error(err)
		return
	}
	// 分发token
	token, err := utils.GenerateToken(int64(u.ID), u.UserName)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	// 返回useInfo_token
	userInfoRes := models.UserLoginResponse{
		Id:        int64(u.ID),
		UserName:  u.UserName,
		CreatedAt: u.CreatedAt.Unix(),
	}
	userResp := models.TokenData{
		AccessToken: token,
		User:        userInfoRes,
	}
	return ctl.RespSuccessWithData(userResp), nil
}
