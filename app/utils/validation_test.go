package utils

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

type validateTestCase struct {
	Value    string
	Expected bool
}

func TestValidate_IsValidEmail(t *testing.T) {
	testCase := []validateTestCase{
		{
			Value:    "lichcse@gmail.com",
			Expected: true,
		},
		{
			Value:    "lichcse @gmail.com",
			Expected: false,
		},
		{
			Value:    "lichcse_gmail.com",
			Expected: false,
		},
		{
			Value:    "lichcse@gmail.com.123.111",
			Expected: true,
		},
	}

	validation := NewIMValidation()
	for _, item := range testCase {
		Equal(t, item.Expected, validation.IsValidEmail(item.Value))
	}
}

func TestValidate_IsValidUsername(t *testing.T) {
	testCase := []validateTestCase{
		{
			Value:    "lichcse",
			Expected: true,
		},
		{
			Value:    "lich cse",
			Expected: false,
		},
		{
			Value:    "lich@cse",
			Expected: false,
		},
		{
			Value:    "lich_tv",
			Expected: false,
		},
		{
			Value:    "lichtv123",
			Expected: true,
		},
	}

	validation := NewIMValidation()
	for _, item := range testCase {
		Equal(t, item.Expected, validation.IsValidUsername(item.Value))
	}
}
