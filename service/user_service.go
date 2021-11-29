// Package service
// created by lilei at 2021/11/26
package service

import (
	"context"
	"iris-blog-server/db"
	"iris-blog-server/ent"
	"iris-blog-server/ent/user"
)

type userService struct {
	*ent.UserClient
}

var UserService = &userService{db.PgClient().User}

func (u userService) GetUserByID(id uint) (*ent.User, error) {
	return u.Query().Where(user.IDEQ(int(id)), user.DeleteTimeIsNil()).First(context.TODO())
}
