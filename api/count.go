package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) countPingEndpoint(ctx *gin.Context) {
	count, err := s.redisDatabase.PFCount(ctx, "/api/v1/ping").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
	ctx.JSON(http.StatusOK, gin.H{
		"flag":  0,
		"count": count,
	})
}
