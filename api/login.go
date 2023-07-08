package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) loginEndpoint(ctx *gin.Context) {

	cookie, err := ctx.Cookie("session")
	if err == http.ErrNoCookie {
		var loginRequest LoginRequest
		err = ctx.ShouldBind(&loginRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error when parse post data",
			})
			return
		}

		if !checkUserLogin(loginRequest.Username, loginRequest.Password) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "username or password incorrect",
			})
			return
		}

		exist, err := s.redisDatabase.Exists(ctx, loginRequest.Username).Result()
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "check exist redis error",
			})
			return
		}
		if exist == 0 {
			ok, err := s.redisDatabase.Set(ctx, loginRequest.Username, 1, 10*time.Second).Result()
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "redis fail",
				})
				return
			}
			fmt.Println(ok)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "fail",
		})
		return
	}
	result, err := s.redisDatabase.Get(ctx, cookie).Result()
	if err == redis.Nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "need login again",
		})
		return
	}
	if result == "-1" {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "need login again",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "logined",
		"result":  result,
	})
}
