package service

import (
	"errors"
	"gin_study/factory"
	"gin_study/gen/models"
	"gin_study/gen/query"
	"gin_study/gen/request"
)

type UserService struct{}

func (u UserService) Login(req *request.LoginRequest) (string, error) {
	user, err := query.User.Where(query.User.Name.Eq(req.Name)).First()
	if err != nil {
		return "", errors.New("user and password are incorrect")
	}
	reqPass := factory.Md5Hash(req.Password)
	if reqPass != user.Password {
		return "", errors.New("user and password are incorrect")
	}
	// get token
	token, e := factory.CreateToken(user.Name, user.ID)
	return token, e
}

func (u UserService) Register(req *request.RegisterRequest) (int64, error) {
	userExists, _ := query.User.Where(query.User.Name.Eq(req.Name)).First()
	if userExists != nil {
		return 0, errors.New("user already exists")
	}
	user := models.User{
		Name:     req.Name,
		Password: factory.Md5Hash(req.Password),
	}
	tx := query.Q.Begin()
	err := query.User.Save(&user)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return 0, e
		}
		return 0, err
	}
	if e := tx.Commit(); e != nil {
		return 0, e
	}
	return user.ID, err
}
