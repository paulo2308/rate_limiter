package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"

	"rate_limiter/config"
	"rate_limiter/limiter"
	"rate_limiter/middleware"
)

func StartServer(cfg *config.Config) {
	r := mux.NewRouter()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	strategy := limiter.NewRedisLimiter(redisClient)

	limiterService := limiter.NewLimiter(strategy, limiter.Config{
		RateLimitIP:    cfg.RateLimitIP,
		RateLimitToken: cfg.RateLimitToken,
		BlockDuration:  time.Duration(cfg.BlockTimeSeconds) * time.Second,
	})

	r.Use(middleware.RateLimiterMiddleware(limiterService))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Request OK"))
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
