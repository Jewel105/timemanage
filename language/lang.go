package language

import (
	"embed"
	"fmt"
	"gin_study/logger"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

//go:embed *.toml
var localeFS embed.FS

var bundle *i18n.Bundle
var localeList map[string]*i18n.Localizer

func InitI18n(paths *[]string) {
	bundle = i18n.NewBundle(language.English) // 设置默认语言为英文
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	localeList = make(map[string]*i18n.Localizer)
	for _, lang := range *paths {
		fmt.Printf("Load language: %s\n", lang)
		_, err := bundle.LoadMessageFileFS(localeFS, fmt.Sprintf("%s.toml", lang))
		if err != nil {
			logger.Error(
				zap.Any("LoadMessageFileFS", err),
			)
		}
		localeList[lang] = i18n.NewLocalizer(bundle, lang, "en")
	}
}

func GetLocale(lang string, id string) string {
	localizer := localeList[lang]
	if localizer == nil {
		return ""
	}
	message, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: id,
	})
	if err != nil {
		logger.Error(
			zap.Any("localizer.Localize", err),
		)
		return ""
	}
	return message
}
