package main

import (
	"fmt"
	"github.com/goget-milk/sso/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// TODO: Init logger
	// TODO: Init application (app)
	// TODO: Run applications gRPC-server
}
