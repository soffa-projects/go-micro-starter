package locales

import (
	"embed"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

//go:embed *.toml
var localesFS embed.FS

var localizer *i18n.Localizer

func Init(defaultLocale string, otherLocales ...string) {
	//localizer = i18n.NewLocalizer(env.Bundle, "fr")

	bundle := i18n.NewBundle(language.French)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err := bundle.LoadMessageFileFS(localesFS, fmt.Sprintf("locale.%s.toml", defaultLocale))
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, lang := range otherLocales {
		_, err := bundle.LoadMessageFileFS(localesFS, fmt.Sprintf("locale.%s.toml", lang))
		if err != nil {
			panic(err)
		}
	}
	allLocales := append(otherLocales, defaultLocale)
	localizer = i18n.NewLocalizer(bundle, allLocales...)
	log.Infof("%d locales loaded", len(allLocales))
}

func T(messageId string, other ...string) string {
	theOrder := ""
	if len(other) > 0 {
		theOrder = other[0]
	}
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    messageId,
			Other: theOrder,
		},
	})
}
