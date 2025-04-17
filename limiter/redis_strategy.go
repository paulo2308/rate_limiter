package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisLimiter struct {
	Client *redis.Client
}

func NewRedisLimiter(client *redis.Client) *RedisLimiter {
	return &RedisLimiter{Client: client}
}

func (r *RedisLimiter) IncrementAndCheck(key string, limit int, blockDuration time.Duration) (bool, int, error) {
	ctx := context.Background()

	pipe := r.Client.TxPipeline()
	count := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, time.Second)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, 0, err
	}

	blocked, err := r.Client.Get(ctx, "block:"+key).Result()
	if err == nil && blocked == "true" {
		return false, int(count.Val()), nil
	}

	if count.Val() > int64(limit) {
		r.Client.Set(ctx, "block:"+key, "true", blockDuration)
		return false, int(count.Val()), nil
	}

	return true, int(count.Val()), nil
}
