package utils

import "regexp"

// IMValidation interface
type IMValidation interface {
	IsValidEmail(str string) bool
	IsValidUsername(str string) bool
}

type imValidate struct{}

// NewIMValidation func
func NewIMValidation() IMValidation {
	return &imValidate{}
}

// IsValidEmail func
func (v *imValidate) IsValidEmail(str string) bool {
	var validID = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return validID.MatchString(str)
}

// IsValidUsername func
func (v *imValidate) IsValidUsername(str string) bool {
	var validID = regexp.MustCompile("^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$")
	return validID.MatchString(str)
}
