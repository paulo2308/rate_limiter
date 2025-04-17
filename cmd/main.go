package main

import (
	"rate_limiter/config"
	"rate_limiter/server"
)

func main() {
	cfg := config.LoadConfig()
	server.StartServer(cfg)
}
