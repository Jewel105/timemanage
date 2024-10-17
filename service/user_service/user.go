package userservice

import (
	"gin_study/api/consts"
	"gin_study/config"
	"gin_study/factory"
	"gin_study/gen/models"
	"gin_study/gen/mysql"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"net/smtp"
	"strconv"
	"time"
)

func Login(req *request.LoginRequest) (string, error) {
	user, err := query.User.Where(query.User.Name.Eq(req.Name)).First()
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
	userExists, _ := query.User.Where(query.User.Name.Eq(req.Name)).First()
	if userExists != nil {
		return 0,
			&consts.ApiErr{Code: consts.BAD_REQUEST, Msg: "user already exists"}
	}
	user := models.User{
		Name:     req.Name,
		Password: factory.Md5Hash(req.Password),
	}
	tx := query.Q.Begin()
	err := query.User.Save(&user)
	err = mysql.DeferTx(tx, err)
	return user.ID, err
}

func SendMail(to *request.SendCodeRequest) error {
	key := config.Config.EmailSmpt.RedisKey + to.Email

	// 避免验证码过期
	code, err := factory.RedisGet(key)
	if err == nil || len(code) == 6 {
		return &consts.ApiErr{Code: consts.TOO_FREQUENTLY, Msg: "The email code have already been sent."}
	}

	from := config.Config.EmailSmpt.Email
	password := config.Config.EmailSmpt.Password

	// 发送给多个收件人
	recipients := []string{to.Email}

	// SMTP服务器配置
	smtpHost := config.Config.EmailSmpt.Host
	smtpPort := strconv.Itoa(config.Config.EmailSmpt.Port)

	randomCode := factory.GenerateRandomString(6)

	// 邮件内容：包含标题和正文
	message := []byte("Subject: Code" + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" + "TimeManage Support System:" +
		"\r\n" + randomCode + "\r\n")

	// 认证信息
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// 发送邮件
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, recipients, message)
	if err != nil {
		return err
	}
	err = factory.RedisSet(key, randomCode, time.Minute*2)
	if err != nil {
		return err
	}
	return nil
}
