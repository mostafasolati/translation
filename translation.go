package translation

import (
	"github.com/mostafasolati/translation/english"
	"github.com/mostafasolati/translation/persian"
)

func TranslateWeekDay(lang string, day int) string {
	switch lang {
	case "fa":
		return persian.TranslateWeekDay(day)
	case "en":
		return english.TranslateWeekDay(day)
	default:
		return "نامشخصی"
	}
}
