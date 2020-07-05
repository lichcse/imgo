package utils

import (
	"errors"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestIMConvert_Object(t *testing.T) {
	imConvert := NewIMConvert()

	err := imConvert.Object(Nil, Nil)
	NotEqual(t, nil, err)

	type source struct {
		Field01 string
		Field02 int
	}
	type destination struct {
		Field01 string
		Field03 interface{}
	}
	sourceTest := source{
		Field01: "test",
	}
	destinationTest := destination{}
	err = imConvert.Object(sourceTest, &destinationTest)
	Equal(t, nil, err)
	Equal(t, sourceTest.Field01, destinationTest.Field01)
}

func TestIMConvert_DatabaseError(t *testing.T) {
	imConvert := NewIMConvert()

	err := imConvert.DatabaseError(nil)
	Equal(t, errors.New("undefined"), err)

	err = imConvert.DatabaseError(errors.New("sql: no rows in result set"))
	Equal(t, errors.New("not_found"), err)

	err = imConvert.DatabaseError(errors.New("undefined xyz"))
	Equal(t, errors.New("undefined"), err)
}
