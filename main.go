package main

import (
	"fmt"
	"go-api/infra/config"

	_ "go-api/docs"
)

// @title           Albums API
// @version         1.0
// @description     API para gerenciamento de álbums em Clean Architecture
func main() {
	env := config.LoadEnv()

	router := config.InitializeServer()

	address := fmt.Sprintf("%s:%s", env.Host, env.Port)
	router.Run(address)
}
