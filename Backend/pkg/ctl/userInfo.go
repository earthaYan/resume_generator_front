package ctl

import (
	"context"
	"errors"
)

type key int

var userkey key

type UserInfo struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContex(ctx)
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	return user, nil
}
func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userkey, u)
}
func FromContex(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userkey).(*UserInfo)
	return u, ok
}
