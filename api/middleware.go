package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *Server) increaseCountBlock() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pingRequest PingRequest
		err := ctx.ShouldBind(&pingRequest)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = s.redisDatabase.PFAdd(context.Background(), ctx.FullPath(), pingRequest.Username).Err()
		fmt.Println(ctx.FullPath())
		if err != nil {

		}
		ctx.Next()
	}

}

func (s *Server) increaseCountRanking() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pingRequest PingRequest
		err := ctx.ShouldBind(&pingRequest)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = s.redisDatabase.ZIncrBy(context.Background(), "visitor", 1, fmt.Sprint(pingRequest.Username)).Err()
		if err != nil {

		}
		ctx.Next()
	}
}
