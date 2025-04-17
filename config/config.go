package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr        string
	RedisPassword    string
	RedisDB          int
	RateLimitIP      int
	RateLimitToken   int
	BlockTimeSeconds int
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rateLimitIP, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_IP"))
	rateLimitToken, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_TOKEN"))
	blockTime, _ := strconv.Atoi(os.Getenv("BLOCK_TIME_SECONDS"))

	return &Config{
		RedisAddr:        os.Getenv("REDIS_ADDR"),
		RedisPassword:    os.Getenv("REDIS_PASSWORD"),
		RedisDB:          redisDB,
		RateLimitIP:      rateLimitIP,
		RateLimitToken:   rateLimitToken,
		BlockTimeSeconds: blockTime,
	}
}
