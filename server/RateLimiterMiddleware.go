package server

import (
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var requestLimitMap map[string]int
var mu sync.Mutex

func RateLimitingMiddleware() gin.HandlerFunc {
	requestLimitMap = make(map[string]int)
	return func(c *gin.Context) {
		maxRequests := viper.GetInt("rateLimiter.maxConcurrentRequests")
		clientIP := getClientIP(c.Request)
		mu.Lock()
		if requestLimitMap[clientIP] < maxRequests {
			requestLimitMap[clientIP]++
			mu.Unlock()
			c.Next()

			mu.Lock()
			requestLimitMap[clientIP]--
			mu.Unlock()
		} else {
			mu.Unlock()
			c.AbortWithStatus(http.StatusTooManyRequests)
		}
	}
}

func getClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}

	return ip
}
