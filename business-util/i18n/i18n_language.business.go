package i18nutil

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var _ = language.Chinese
var _ = language.English
var _ = i18n.Message{}

type Language string

const (
	LANG_CHINESES Language = "zh"
	LANG_ENGLISH  Language = "en"
)
