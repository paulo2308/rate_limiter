package tests

import (
	"github.com/go-redis/redis/v8"
	"rate_limiter/limiter"
	"testing"
	"time"
)

func TestRateLimitExceeded(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	rl := limiter.NewRedisLimiter(client)

	key := "test_ip"
	limit := 3
	blockTime := 60

	blockDuration := time.Duration(blockTime) * time.Second

	for i := 0; i < limit; i++ {
		allowed, _, err := rl.IncrementAndCheck(key, limit, blockDuration)
		if err != nil {
			t.Errorf("Error checking rate limit on request %d: %v", i+1, err)
		}
		if !allowed {
			t.Errorf("Request %d was blocked unexpectedly", i+1)
		}
	}

	allowed, _, err := rl.IncrementAndCheck(key, limit, blockDuration)
	if err != nil {
		t.Errorf("Error checking rate limit after limit exceeded: %v", err)
	}
	if allowed {
		t.Error("Expected request to be blocked, but it was allowed")
	}

	time.Sleep(time.Duration(blockTime+2) * time.Second)

	allowed, _, err = rl.IncrementAndCheck(key, limit, blockDuration)
	if err != nil {
		t.Errorf("Error checking rate limit after block period: %v", err)
	}
	if !allowed {
		t.Error("Expected request to be allowed after block time, but it was blocked")
	}
}
