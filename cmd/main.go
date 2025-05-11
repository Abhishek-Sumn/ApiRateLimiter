package main

import (
	"fmt"
	"log"
	"net/http"
	"ratelimiter/redis"
	"ratelimiter/routes"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env vars")
	}
	redis.InitRedis()

	r := routes.SetupRouter()
	fmt.Println("Server booming on 9090")
	http.ListenAndServe(":9090", r)
}
