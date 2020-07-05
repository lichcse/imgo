package utils

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestIMCrypt_Hash_CheckHash(t *testing.T) {
	str := "password"
	imCrypt := NewIMCrypt()
	hash, err := imCrypt.Hash(str)
	Equal(t, nil, err)
	Equal(t, true, imCrypt.CheckHash(str, hash))
}
