package service

import "github.com/remoting/frame/server/auth"

type UserService struct {
}

func (*UserService) GetUserInfoById(userId string) (auth.UserInfo, error) {
	return nil, nil
}
