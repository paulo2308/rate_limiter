package limiter

import "time"

type LimiterStrategy interface {
	IncrementAndCheck(key string, limit int, blockDuration time.Duration) (allowed bool, count int, err error)
}
