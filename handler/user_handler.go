// Package handler
// created by lilei at 2021/11/23
package handler

import (
	"github.com/kataras/iris/v12"
	"iris-blog-server/config"
	"iris-blog-server/service"
)

type UserHandler struct{ BaseHandler }

func (u *UserHandler) GetBy(id uint) Response {
	user, err := service.UserService.GetUserByID(id)
	if err != nil {
		return u.Fail(err)
	}
	if user == nil {
		return u.Fail(config.UserNotExistsError)
	}
	return u.SuccessData(iris.Map{"user": user})
}
