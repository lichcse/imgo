package utils

import "time"

// IMTime interface of time object
type IMTime interface {
	TimeDB() string
}

type imTime struct{}

// NewIMTime func new time object
func NewIMTime() IMTime {
	return &imTime{}
}

// TimeDB func get current time with database format
func (t *imTime) TimeDB() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}
