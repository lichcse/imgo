package utils

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestBigCache_Get(t *testing.T) {
	bigCache := NewBigCache()
	err := bigCache.Set("test_key_big_cache", "test_value_big_cache")
	Equal(t, nil, err)
	value, err := bigCache.Get("test_key_big_cache")
	Equal(t, nil, err)
	Equal(t, "test_value_big_cache", value)
	value, err = bigCache.Get("test_key_big_cache_fail")
	NotEqual(t, nil, err)
	Equal(t, "", value)
}
