package api

import (
	"net/http"
	"time"

	"github.com/bsm/redislock"
	"github.com/gin-gonic/gin"
)

type PingRequest struct {
	Username string `json:"username"`
}

func (s *Server) pingEndpoint(ctx *gin.Context) {
	redisLock := redislock.New(s.redisDatabase)
	lock, err := redisLock.Obtain(ctx, "api:ping", 10*time.Second, nil)
	if err == redislock.ErrNotObtained {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "only 1 person can ping at a time",
		})
		return
	}
	defer lock.Release(ctx)

	time.Sleep(5 * time.Second)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
