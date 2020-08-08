package utils

import (
	"encoding/json"
	"errors"
)

// IMConvert interface
type IMConvert interface {
	Object(source interface{}, destination interface{}) error
	DatabaseError(err error) error
}

type imConvert struct{}

// NewIMConvert func
func NewIMConvert() IMConvert {
	return &imConvert{}
}

// Object func
func (c *imConvert) Object(source interface{}, destination interface{}) error {
	byteConvert, err := json.Marshal(source)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteConvert, &destination)
	return err
}

// DatabaseError func
func (c *imConvert) DatabaseError(err error) error {
	if err == nil {
		return errors.New("undefined")
	}

	switch err.Error() {
	case "sql: no rows in result set":
		return errors.New("not_found")
	default:
		return errors.New("undefined")
	}
}
