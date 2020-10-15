package utils

import (
	"testing"
	//"github.com/gomodule/redigo/redis"
	//. "github.com/stretchr/testify/assert"
)

func initRedisCache() RedisCache {
	return NewRedis(RedisConfig{
		URL: "redis://127.0.0.1:6379/0",
	})
}

func TestRedisCache_Conn(t *testing.T) {
	// redisCache := initRedisCache()
	// c := redisCache.Conn()
	// defer c.Close()
	// c.Send("MULTI")
	// c.Send("INCR", "foo")
	// c.Send("INCR", "bar")
	// res, err := redis.Ints(c.Do("EXEC"))
	// Equal(t, nil, err)
	// Equal(t, 2, len(res))
}

func TestRedisCache_Keys(t *testing.T) {
	// redisCache := initRedisCache()
	// err := redisCache.Set("test_set_get_key", "test_set_get_key", 1)
	// Equal(t, nil, err)
	// keys, err := redisCache.Keys("*")
	// Equal(t, nil, err)
	// LessOrEqual(t, 1, len(keys))
}

func TestRedisCache_Set(t *testing.T) {
	// redisCache := initRedisCache()
	// err := redisCache.Set("test_set_get_key", "test_set_get_key", 1)
	// Equal(t, nil, err)
	// val, err := redisCache.Get("test_set_get_key")
	// Equal(t, nil, err)
	// Equal(t, "test_set_get_key", val)
	// err = redisCache.Set("test_set_get_key", "test_set_get_key", -1)
	// Equal(t, nil, err)
}

func TestRedisCache_Get(t *testing.T) {
	// redisCache := initRedisCache()
	// err := redisCache.Set("test_set_get_key", "test_set_get_key", 1)
	// Equal(t, nil, err)
	// val, err := redisCache.Get("test_set_get_key")
	// Equal(t, nil, err)
	// Equal(t, "test_set_get_key", val)
}
