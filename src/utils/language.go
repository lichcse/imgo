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
func (m *imLanguage) SetLanguage(lang string) {
	m.currentLang = lang
}

// SetUnknownMessage func
// The unknown message will return when it not found in the current or default language.
func (m *imLanguage) SetUnknownMessage(message string) {
	m.unknownMessage = message
}

// GetMessage func
// This function will return the message with the corresponding code.
func (m *imLanguage) GetMessage(code string) string {
	lang, useCurrentLang := m.dataLang[m.currentLang]
	if !useCurrentLang {
		lang = m.dataLang[m.defaultLang]
	}

	if message, ok := lang[code]; ok {
		return message
	}

	if !useCurrentLang {
		return m.unknownMessage
	}

	if lang, ok := m.dataLang[m.defaultLang]; ok {
		if message, ok := lang[code]; ok {
			return message
		}
	}
	return m.unknownMessage
}
