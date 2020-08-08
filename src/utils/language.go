package utils

// Lang struct
type Lang map[string]string

// IMLanguage interface
type IMLanguage interface {
	GetMessage(code string) string
	SetLanguage(lang string)
	SetUnknownMessage(message string)
}

type imLanguage struct {
	currentLang    string
	defaultLang    string
	unknownMessage string
	dataLang       map[string]Lang
}

// NewIMLanguage func
func NewIMLanguage(dataLang map[string]Lang, defaultLang string) IMLanguage {
	return &imLanguage{dataLang: dataLang, defaultLang: defaultLang, unknownMessage: "Unknown message."}
}

// SetLanguage func
func (l *imLanguage) SetLanguage(lang string) {
	l.currentLang = lang
}

// SetUnknownMessage func
// The unknown message will return when it not found in the current or default language.
func (l *imLanguage) SetUnknownMessage(message string) {
	l.unknownMessage = message
}

// GetMessage func
// This function will return the message with the corresponding code.
func (l *imLanguage) GetMessage(code string) string {
	lang, useCurrentLang := l.dataLang[l.currentLang]
	if !useCurrentLang {
		lang = l.dataLang[l.defaultLang]
	}

	if message, ok := lang[code]; ok {
		return message
	}

	if !useCurrentLang {
		return l.unknownMessage
	}

	if lang, ok := l.dataLang[l.defaultLang]; ok {
		if message, ok := lang[code]; ok {
			return message
		}
	}
	return l.unknownMessage
}
