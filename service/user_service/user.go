package userservice

import (
	"gin_study/api/consts"
	"gin_study/config"
	"gin_study/factory"
	"gin_study/gen/models"
	"gin_study/gen/mysql"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"gin_study/logger"
	"net/smtp"
	"strconv"
	"time"

	"go.uber.org/zap"
)

func Login(req *request.LoginRequest) (string, error) {
	// 用户名或密码都可以登录
	user, err := query.User.Where(query.User.Name.Eq(req.Name)).Or(query.User.Email.Eq(req.Name)).First()
	if err != nil {
		return "", &consts.ApiErr{Code: consts.LOGIN_FAILED, Msg: "user and password are incorrect"}
	}
	reqPass := factory.Md5Hash(req.Password)
	if reqPass != user.Password {
		return "", &consts.ApiErr{Code: consts.LOGIN_FAILED, Msg: "user and password are incorrect"}
	}
	// get token
	token, e := factory.CreateToken(user.Name, user.ID)
	return token, e
}

func Register(req *request.RegisterRequest) (int64, error) {
	nameExists, _ := query.User.Select(query.User.ID).Where(query.User.Name.Eq(req.Name)).First()
	if nameExists != nil {
		return 0,
			&consts.ApiErr{Code: consts.BAD_REQUEST, Msg: "name already exists."}
	}
	emailExists, _ := query.User.Select(query.User.ID).Where(query.User.Email.Eq(req.Email)).First()
	if emailExists != nil {
		return 0,
			&consts.ApiErr{Code: consts.BAD_REQUEST, Msg: "email already exists."}
	}
	// 获取验证码并检查
	key := config.Config.EmailSmpt.RedisKey + req.Email
	code, err := factory.RedisGet(key)
	if err != nil || code != req.Code {
		return 0,
			&consts.ApiErr{Code: consts.CODE_INVALID, Msg: "email code is invalid."}
	}
	err = factory.RedisDel(key)
	if err != nil {
		return 0, err
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: factory.Md5Hash(req.Password),
	}
	tx := query.Q.Begin()
	err = query.User.Save(&user)
	if err != nil {
		return 0, err
	}
	err = mysql.DeferTx(tx, err)
	return user.ID, err
}

func SendCode(to *request.SendCodeRequest) error {
	key := config.Config.EmailSmpt.RedisKey + to.Email

	// 避免验证码过于频繁
	code, err := factory.RedisGet(key)
	if err == nil || len(code) == 6 {
		return &consts.ApiErr{Code: consts.TOO_FREQUENTLY, Msg: "The email code have already been sent."}
	}
	randomCode := factory.GenerateRandomString(6)

	// 线上环境才真正发送邮箱
	if config.Env == "pro" {
		err = sendMailOnline(to.Email, randomCode)
		if err != nil {
			return err
		}
	}

	err = factory.RedisSet(key, randomCode, time.Minute*2)
	if err != nil {
		return err
	}
	logger.Info("[SendMail]",
		zap.String("email", to.Email),
		zap.String("code", randomCode),
	)
	return nil
}

func sendMailOnline(email string, code string) error {
	from := config.Config.EmailSmpt.Email
	password := config.Config.EmailSmpt.Password

	// 发送给多个收件人
	recipients := []string{email}
	// SMTP服务器配置
	smtpHost := config.Config.EmailSmpt.Host
	smtpPort := strconv.Itoa(config.Config.EmailSmpt.Port)

	// 邮件内容：包含标题和正文
	message := []byte("Subject: Code" + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" + "TimeManage Support System:" +
		"\r\n" + code + "\r\n")

	// 认证信息
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// 发送邮件
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, recipients, message)
	return err
}

func ForgetPassword(req *request.RegisterRequest) (int64, error) {
	user, err := query.User.Where(query.User.Name.Eq(req.Name), query.User.Email.Eq(req.Email)).First()
	if err != nil {
		return 0,
			&consts.ApiErr{Code: consts.NO_DATA, Msg: "user not found."}
	}

	if user.Password == factory.Md5Hash(req.Password) {
		return 0,
			&consts.ApiErr{Code: consts.BAD_REQUEST, Msg: "The new password is the same as the old password."}
	}

	// 获取验证码并检查
	key := config.Config.EmailSmpt.RedisKey + req.Email
	code, err := factory.RedisGet(key)
	if err != nil || code != req.Code {
		return 0,
			&consts.ApiErr{Code: consts.CODE_INVALID, Msg: "email code is invalid."}
	}
	err = factory.RedisDel(key)
	if err != nil {
		return 0, err
	}

	user.Password = factory.Md5Hash(req.Password)

	tx := query.Q.Begin()
	err = query.User.Save(user)
	if err != nil {
		return 0, err
	}
	err = mysql.DeferTx(tx, err)
	return user.ID, err
}
