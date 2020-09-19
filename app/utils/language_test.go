package utils

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

var unknownMessage = "Unknown message."

func initDataLang() map[string]Lang {
	return map[string]Lang{
		"en": map[string]string{
			"test_01": "en_test_01",
			"test_02": "en_test_02",
			"test_04": "en_test_04",
		},
		"vi": map[string]string{
			"test_01": "vi_test_01",
			"test_03": "vi_test_03",
			"test_04": "vi_test_04",
		},
	}
}

func TestLanguage_UnknownMessage(t *testing.T) {
	dataLang := initDataLang()
	lang := NewIMLanguage(dataLang, "en")
	Equal(t, unknownMessage, lang.GetMessage(""))
	Equal(t, unknownMessage, lang.GetMessage("test_03"))
	Equal(t, unknownMessage, lang.GetMessage("test_0x"))
	lang.SetUnknownMessage("New unknown message.")
	Equal(t, "New unknown message.", lang.GetMessage(""))
}

func TestLanguage_DefaultLang(t *testing.T) {
	dataLang := initDataLang()
	lang := NewIMLanguage(dataLang, "vi")
	Equal(t, "vi_test_01", lang.GetMessage("test_01"))
	lang.SetLanguage("en")
	Equal(t, "en_test_02", lang.GetMessage("test_02"))
	lang = NewIMLanguage(dataLang, "unknown")
	Equal(t, unknownMessage, lang.GetMessage("test_04"))
}

func TestLanguage_GetMessage(t *testing.T) {
	dataLang := initDataLang()
	lang := NewIMLanguage(dataLang, "en")
	lang.SetLanguage("ru")
	Equal(t, "en_test_01", lang.GetMessage("test_01"))
	Equal(t, unknownMessage, lang.GetMessage("test_03"))
	lang.SetLanguage("vi")
	Equal(t, unknownMessage, lang.GetMessage("test_0x"))
	Equal(t, "vi_test_01", lang.GetMessage("test_01"))
	Equal(t, "en_test_02", lang.GetMessage("test_02"))
}
