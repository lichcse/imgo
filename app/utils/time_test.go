package utils

import (
	"testing"
	"time"

	. "github.com/stretchr/testify/assert"
)

func TestTime_TimeDB(t *testing.T) {
	imTime := NewIMTime()
	timeTest := time.Date(
		2020, 06, 01, 0, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05")
	Greater(t, imTime.TimeDB(), timeTest)
}
