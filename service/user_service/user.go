package userservice

import (
	"gin_study/api/consts"
	"gin_study/factory"
	"gin_study/gen/models"
	"gin_study/gen/query"
	"gin_study/gen/request"
)

func Login(req *request.LoginRequest) (string, error) {
	user, err := query.User.Where(query.User.Name.Eq(req.Name)).First()
	if err != nil {
		return "", consts.ApiErr{Code: consts.LOGIN_FAILED, Msg: "user and password are incorrect"}
	}
	reqPass := factory.Md5Hash(req.Password)
	if reqPass != user.Password {
		return "", consts.ApiErr{Code: consts.LOGIN_FAILED, Msg: "user and password are incorrect"}
	}
	// get token
	token, e := factory.CreateToken(user.Name, user.ID)
	return token, e
}

func Register(req *request.RegisterRequest) (int64, error) {
	userExists, _ := query.User.Where(query.User.Name.Eq(req.Name)).First()
	if userExists != nil {
		return 0,
			consts.ApiErr{Code: consts.BAD_REQUEST, Msg: "user already exists"}
	}
	user := models.User{
		Name:     req.Name,
		Password: factory.Md5Hash(req.Password),
	}
	tx := query.Q.Begin()
	err := query.User.Save(&user)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return user.ID, err
}
