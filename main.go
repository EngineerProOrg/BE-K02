package main

import (
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/thnam4500/identity/api"
)

type PingRequest struct {
	Username int `json:"user_name"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var rdb *redis.Client

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()
	s := api.NewServer(rdb)
	err := s.StartServer(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
