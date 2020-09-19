package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_TableName(t *testing.T) {
	user := User{}
	table := user.TableName()
	assert.Equal(t, "im_user", table)
}
