package cache

import (
	"imgo/src/utils"
	"time"

	// redis
	"github.com/gomodule/redigo/redis"
)

// RedisCache interface
type RedisCache interface {
	Conn() redis.Conn
	Keys(key string) ([]string, error)
	Set(key string, value string, expired int64) error
	Get(key string) (string, error)
}

type redisCache struct {
	config utils.RedisConfig
	pool   *redis.Pool
}

// NewRedis func
func NewRedis(config utils.RedisConfig) RedisCache {
	p := &redis.Pool{
		MaxIdle:     1000,
		MaxActive:   1000,
		IdleTimeout: 1 * time.Minute,
		Wait:        true,
	}
	rc := &redisCache{config: config, pool: p}
	p.Dial = rc.dialURL
	return rc
}

// Conn func
// Please note when you using this function you need to close connections
func (m *redisCache) Conn() redis.Conn {
	return m.pool.Get()
}

// dialURL func
func (m *redisCache) dialURL() (redis.Conn, error) {
	return redis.DialURL(m.config.URL)
}

// Keys func
func (m *redisCache) Keys(pattern string) ([]string, error) {
	conn := m.Conn()
	defer conn.Close()
	return redis.Strings(conn.Do("KEYS", pattern))
}

// Set func
func (m *redisCache) Set(key string, value string, expired int64) error {
	conn := m.Conn()
	defer conn.Close()

	var err error
	if expired < 0 {
		_, err = conn.Do("SET", key, value)
	} else {
		_, err = conn.Do("SETEX", key, expired, value)
	}
	return err
}

// Get func
func (m *redisCache) Get(key string) (string, error) {
	conn := m.Conn()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}
