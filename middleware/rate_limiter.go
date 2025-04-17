package middleware

import (
	"net/http"
	"rate_limiter/limiter"
	"strings"
)

func RateLimiterMiddleware(service *limiter.LimiterService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("API_KEY")
			ip := strings.Split(r.RemoteAddr, ":")[0]

			allowed, err := service.Allow(ip, token)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			if !allowed {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("you have reached the maximum number of requests or actions allowed within a certain time frame"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
