package utils

import "time"

// IMTime interface
type IMTime interface {
	TimeDB() string
}

type imTime struct{}

// NewIMTime func
func NewIMTime() IMTime {
	return &imTime{}
}

// TimeDB func
func (t *imTime) TimeDB() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}
