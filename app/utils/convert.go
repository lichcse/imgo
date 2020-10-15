package utils

import (
	"encoding/json"
	"errors"
)

// IMConvert interface of convert object
type IMConvert interface {
	Object(source interface{}, destination interface{}) error
	DatabaseError(err error) error
}

type imConvert struct{}

// NewIMConvert func new convert object
func NewIMConvert() IMConvert {
	return &imConvert{}
}

// Object func convert object
func (c *imConvert) Object(source interface{}, destination interface{}) error {
	byteConvert, err := json.Marshal(source)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteConvert, &destination)
	return err
}

// DatabaseError func convert database error to app error code
func (c *imConvert) DatabaseError(err error) error {
	if err == nil {
		return errors.New("undefined")
	}

	switch err.Error() {
	case "record not found":
		return errors.New("not_found")
	case "sql: no rows in result set":
		return errors.New("not_found")
	default:
		return errors.New("undefined")
	}
}
