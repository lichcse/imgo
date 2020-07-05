package resource

import (
	"imgo/src/modules/identity/v1/resource/lang"
	"imgo/src/utils"
)

// DefaultLang default language
var DefaultLang = "en"

// IdentityLang identity language
var IdentityLang = map[string]utils.Lang{
	"en": lang.EN,
	"vi": lang.VI,
}
