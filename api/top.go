package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) topEndpoint(ctx *gin.Context) {
	result, err := s.redisDatabase.ZRevRangeWithScores(ctx, "visitor", 0, 10).Result()
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}
