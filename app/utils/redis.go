package utils

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisCache interface of redis object
type RedisCache interface {
	Conn() redis.Conn
	Keys(key string) ([]string, error)
	Set(key string, value string, expired int64) error
	Get(key string) (string, error)
}

type redisCache struct {
	config RedisConfig
	pool   *redis.Pool
}

// NewRedis func new redis object
func NewRedis(config RedisConfig) RedisCache {
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

// Conn func get redis connection
// Please note when you using this function you need to close connections
func (r *redisCache) Conn() redis.Conn {
	return r.pool.Get()
}

func (r *redisCache) dialURL() (redis.Conn, error) {
	return redis.DialURL(r.config.URL)
}

// Keys func get all keys matching pattern
func (r *redisCache) Keys(pattern string) ([]string, error) {
	conn := r.Conn()
	defer conn.Close()
	return redis.Strings(conn.Do("KEYS", pattern))
}

// Set func set key to hold the string value
func (r *redisCache) Set(key string, value string, expired int64) error {
	conn := r.Conn()
	defer conn.Close()

	var err error
	if expired < 0 {
		_, err = conn.Do("SET", key, value)
	} else {
		_, err = conn.Do("SETEX", key, expired, value)
	}
	return err
}

// Get func get the value of key
func (r *redisCache) Get(key string) (string, error) {
	conn := r.Conn()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}
