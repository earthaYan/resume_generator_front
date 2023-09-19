package repositories

import (
	"context"
	"fmt"
	"resume_backend/models"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func (s *UserDao) FindUserByUserName(userName string) (user *models.User, err error) {
	err = s.DB.Model(&models.User{}).Where("user_name=?", userName).First(&user).Error

	return
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (s *UserDao) CreateUser(user *models.User) (err error) {
	err = s.DB.Model(&models.User{}).Create(user).Error
	fmt.Println("error", err)
	return
}
