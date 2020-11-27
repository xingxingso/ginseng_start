package service

import "ginseng_start/model"

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (l *Login) Login() *model.User {
	if l.Username == "admin" && l.Password == "admin" {
		return &model.User{
			Id:   1,
			Name: l.Username,
		}
	}

	return nil
}
