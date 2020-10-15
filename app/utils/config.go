package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// MongoConfig struct mongo config info
type MongoConfig struct {
	URL       string `yaml:"url"`
	Database  string `yaml:"database"`
	PoolLimit int    `yaml:"pool_limit"`
}

// MySQLConfig struct mysql config info
type MySQLConfig struct {
	URL         string `yaml:"url"`
	Database    string `yaml:"database"`
	PoolLimit   int    `yaml:"pool_limit"`
	MaxLifetime int    `yaml:"max_life_time"`
	LogMode     bool   `yaml:"log_mode"`
}

// RedisConfig struct redis config info
type RedisConfig struct {
	URL string `yaml:"url"`
}

// RabbitMQConfig struct rabbit config info
type RabbitMQConfig struct {
	URL string `yaml:"url"`
}

// EnvironmentConfig struct env config info
type EnvironmentConfig struct {
	Mongo    map[string]MongoConfig `yaml:"mongo"`
	MySQL    map[string]MySQLConfig `yaml:"mysql"`
	Redis    RedisConfig            `yaml:"redis"`
	RabbitMQ RabbitMQConfig         `yaml:"rabbitmq"`
	PORT     string                 `yaml:"port"`
}

// IMConfig interface of config object
type IMConfig interface {
	Load(args []string) error
	GetPort() string
	Mongo() map[string]MongoConfig
	MongoItem(db string) MongoConfig
	MySQL() map[string]MySQLConfig
	MySQLItem(db string) MySQLConfig
	RabbitMQ() RabbitMQConfig
}

type imConfig struct {
	projectName string
}

var cfg = EnvironmentConfig{}
var cfgMapping = map[string]string{
	"dev":       "dev.yaml",
	"stg":       "stg.yaml",
	"prd":       "prd.yaml",
	"default":   "dev.yaml",
	"test":      "test.yaml",
	"unit_test": "unit_test.yaml",
}

// NewIMConfig func new config object
func NewIMConfig() IMConfig {
	return &imConfig{projectName: "imgo"}
}

// Load func load data from config file
func (c *imConfig) Load(args []string) error {
	if len(args) < 1 {
		return errors.New("")
	}

	env, ok := cfgMapping[args[0]]
	if !ok {
		env = cfgMapping["default"]
	}

	_, filename, _, _ := runtime.Caller(0)
	ddd := strings.Split(filename, c.projectName)

	configData, err := ioutil.ReadFile(ddd[0] + c.projectName + "/config/" + env)
	if err != nil {
		return err
	}

	cfg = EnvironmentConfig{}
	return yaml.Unmarshal(configData, &cfg)
}

// GetPort func get server port
func (c *imConfig) GetPort() string {
	if len(cfg.PORT) <= 0 {
		return ":8080"
	}

	return fmt.Sprintf(":%s", cfg.PORT)
}

// Mongo func get all mongo config info
func (c *imConfig) Mongo() map[string]MongoConfig {
	return cfg.Mongo
}

// MongoItem func get item mongo config info
func (c *imConfig) MongoItem(db string) MongoConfig {
	if dbInfo, ok := cfg.Mongo[db]; ok {
		return dbInfo
	}

	return MongoConfig{}
}

// MySQL func get all mysql config info
func (c *imConfig) MySQL() map[string]MySQLConfig {
	return cfg.MySQL
}

// MySQLItem func get item mysql config info
func (c *imConfig) MySQLItem(db string) MySQLConfig {
	if dbInfo, ok := cfg.MySQL[db]; ok {
		return dbInfo
	}

	return MySQLConfig{}
}

// RabbitMQ func get rabbit config info
func (c *imConfig) RabbitMQ() RabbitMQConfig {
	return cfg.RabbitMQ
}
