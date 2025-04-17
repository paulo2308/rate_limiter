package limiter

import (
	"fmt"
	"time"
)

type LimiterService struct {
	strategy LimiterStrategy
	config   Config
}

type Config struct {
	RateLimitIP    int
	RateLimitToken int
	BlockDuration  time.Duration
}

func NewLimiter(strategy LimiterStrategy, config Config) *LimiterService {
	return &LimiterService{
		strategy: strategy,
		config:   config,
	}
}

func (l *LimiterService) Allow(ip string, token string) (bool, error) {
	var key string
	var limit int

	if token != "" {
		key = fmt.Sprintf("token:%s", token)
		limit = l.config.RateLimitToken
	} else {
		key = fmt.Sprintf("ip:%s", ip)
		limit = l.config.RateLimitIP
	}

	allowed, count, err := l.strategy.IncrementAndCheck(key, limit, l.config.BlockDuration)
	if err != nil {
		return false, err
	}

	if !allowed {
		fmt.Printf("Bloqueado: %s (count: %d)\n", key, count)
	}
	return allowed, nil
}
