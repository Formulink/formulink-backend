package redis

import "github.com/redis/go-redis/v9"

type RConfig struct {
	Address  string `yaml:"REDIS_ADDRESS" env:"REDIS_ADDRESS" env-default:"localhost:6379"`
	Password string `yaml:"REDIS_PASSWORD" env:"REDIS_PASSWORD" env-default:""`
	DB       int    `yaml:"REDIS_DB" env:"REDIS_DB" env-default:"0"`
	Protocol int    `yaml:"REDIS_PROTOCOL" env:"REDIS_PROTOCOL" env-default:"2"`
}

func NewRedisConn(c RConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       c.DB,
		Protocol: c.Protocol,
	})

	return client
}
