package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"github.com/xingxingso/ginseng_start/model"
)


type Credential struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (s *Credential) Login() (*model.User, error) {
	user, err := s.getUser()
	if err != nil {
		return nil, err
	}

	if !s.verifyPwd(user.Password) {
		logrus.Errorf("error password %s %s", s.Password, user.Password)
		return nil, errors.New("password error")
	}

	return nil, nil
}

func (s *Credential) SignUp() (*model.User, error) {
	_, err := s.getUser()
	if err == nil {
		return nil, errors.New("user name exists")
	}

	hashedPassword, err := s.encryptPwd()
	if err != nil {
		return nil, err
	}

	userService := &User{}
	user, err := userService.createUser(s.Username, hashedPassword)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Credential) getUser() (*model.User, error) {
	userService := &User{}
	user, err := userService.getUserByName(s.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Credential) verifyPwd(encryptPwd string) bool  {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptPwd), []byte(s.Password)); err != nil {
		logrus.Infof("verifyPwd err: %v", err)
		return  false
	}
	return true
}

func (s *Credential) encryptPwd() (string, error)  {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s.Password), 8)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}