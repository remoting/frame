package service

import "github.com/remoting/frame/spring"

var beans = spring.New(nil)

func init() {
	beans.Reg(&UserService{})
}

func GetSpring() spring.Spring {
	return beans
}
