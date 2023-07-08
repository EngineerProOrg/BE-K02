package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	redisDatabase *redis.Client
	router        *gin.Engine
}

func NewServer(rdb *redis.Client) *Server {
	server := &Server{redisDatabase: rdb}
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/ping", server.increaseCountRanking(), server.increaseCountBlock(), server.rateLimitUser(), server.pingEndpoint)
		v1.POST("/login", server.loginEndpoint)
		v1.GET("/top", server.topEndpoint)
		v1.GET("/count", server.countPingEndpoint)
	}

	server.router = router
	return server
}

func (s *Server) StartServer(address string) error {
	return s.router.Run(address)
}

func checkUserLogin(username string, password string) bool {
	if username == "admin" && password == "admin" {
		return true
	}
	return false
}

func (s *Server) rateLimitUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limiter := redis_rate.NewLimiter(s.redisDatabase)
		var pingRequest PingRequest
		res, err := limiter.Allow(ctx, fmt.Sprint(pingRequest.Username), redis_rate.PerMinute(2))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "redis error",
			})
			ctx.Abort()
		}

		if res.Allowed == 0 {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"message": "too many requests",
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
