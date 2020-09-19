package resource

import (
	"imgo/app/modules/identity/v1/resource/lang"
	"imgo/app/utils"
)

// DefaultLang default language
var DefaultLang = "en"

// IdentityLang identity language
var IdentityLang = map[string]utils.Lang{
	"en": lang.EN,
	"vi": lang.VI,
}
