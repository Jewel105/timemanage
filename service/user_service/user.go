package userservice

import (
	"fmt"
	"gin_study/api/consts"
	"gin_study/config"
	"gin_study/factory"
	"gin_study/gen/models"
	"gin_study/gen/mysql"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"gin_study/language"
	"gin_study/logger"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

func Login(equipmentID int64, req *request.LoginRequest, lang string) (string, error) {
	// 用户名或密码都可以登录
	user, err := query.User.Select(query.User.ID, query.User.Name, query.User.Password, query.User.Email).Where(query.User.Name.Eq(req.Name)).Or(query.User.Email.Eq(req.Name)).First()
	if err != nil {
		return "", &consts.ApiErr{Code: consts.LOGIN_FAILED, Msg: language.GetLocale(lang, "UserOrPasswordError")}
	}
	reqPass := factory.Md5Hash(req.Password)
	if reqPass != user.Password {
		return "", &consts.ApiErr{Code: consts.LOGIN_FAILED, Msg: language.GetLocale(lang, "UserOrPasswordError")}
	}

	// 记录设备
	if equipmentID != 0 {
		queryEquipment := query.Equipment.Where(query.Equipment.ID.Eq(equipmentID))
		equipment, _ := queryEquipment.Select(query.Equipment.UserIDs).First()
		if equipment != nil {
			// 用户中已经有该设备，不再记录
			userIDStr := strconv.FormatInt(user.ID, 10)
			if !strings.Contains(equipment.UserIDs, userIDStr) {
				userIDs := fmt.Sprintf("%s,%d", equipment.UserIDs, user.ID)
				tx := query.Q.Begin()
				_, err = queryEquipment.UpdateColumn(query.Equipment.UserIDs, userIDs)
				mysql.DeferTx(tx, err)
			}
		}
	}

	// get token
	token, e := factory.CreateToken(user.Name, user.ID)

	return token, e
}

func Logout(userID int64, tokenID uuid.UUID, lang string) (bool, error) {
	err := factory.DeleteToken(userID, tokenID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Register(req *request.RegisterRequest, lang string) (int64, error) {
	nameExists, _ := query.User.Select(query.User.ID).Where(query.User.Name.Eq(req.Name)).First()
	if nameExists != nil {
		return 0,
			&consts.ApiErr{Code: consts.BAD_REQUEST, Msg: language.GetLocale(lang, "NameExits")}
	}
	emailExists, _ := query.User.Select(query.User.ID).Where(query.User.Email.Eq(req.Email)).First()
	if emailExists != nil {
		return 0,
			&consts.ApiErr{Code: consts.BAD_REQUEST, Msg: language.GetLocale(lang, "EmailExits")}
	}
	// 获取验证码并检查
	key := config.Config.EmailSmpt.RedisKey + req.Email
	code, err := factory.RedisGet(key)
	if err != nil || code != req.Code {
		return 0,
			&consts.ApiErr{Code: consts.CODE_INVALID, Msg: language.GetLocale(lang, "EmailCodeError")}
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

func SendCode(to *request.SendCodeRequest, lang string) error {
	key := config.Config.EmailSmpt.RedisKey + to.Email

	// 避免验证码过于频繁
	code, err := factory.RedisGet(key)
	if err == nil || len(code) == 6 {
		return &consts.ApiErr{Code: consts.TOO_FREQUENTLY, Msg: language.GetLocale(lang, "CodeSent")}
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

func sendMailOnline(toEmail string, code string) error {
	from := config.Config.EmailSmpt.Email
	password := config.Config.EmailSmpt.Password
	// SMTP服务器配置
	smtpHost := config.Config.EmailSmpt.Host
	smtpPort := config.Config.EmailSmpt.Port

	msg := gomail.NewMessage()
	msg.SetHeader("From", "TimeManage <"+from+">")
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", "TimeManage Support")
	msg.SetBody("text/html", "Verification code: "+
		"\r\n"+code+"\r\n")

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)
	err := d.DialAndSend(msg)
	return err
}

func ForgetPassword(req *request.RegisterRequest, lang string) (int64, error) {
	user, err := query.User.Where(query.User.Name.Eq(req.Name), query.User.Email.Eq(req.Email)).First()
	if err != nil {
		return 0,
			&consts.ApiErr{Code: consts.NO_DATA, Msg: language.GetLocale(lang, "UserOrEmailError")}
	}

	if user.Password == factory.Md5Hash(req.Password) {
		return 0,
			&consts.ApiErr{Code: consts.BAD_REQUEST, Msg: language.GetLocale(lang, "PasswordEqual")}
	}

	// 获取验证码并检查
	key := config.Config.EmailSmpt.RedisKey + req.Email
	code, err := factory.RedisGet(key)
	if err != nil || code != req.Code {
		return 0,
			&consts.ApiErr{Code: consts.CODE_INVALID, Msg: language.GetLocale(lang, "EmailInvalid")}
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
