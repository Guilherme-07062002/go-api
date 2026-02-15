package main

import (
	"fmt"
	"go-api/infra/config"

	_ "go-api/docs"
)

// @title           Albums API
// @version         1.0
// @description     API para gerenciamento de álbums em Clean Architecture

// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Digite: Bearer {token}
func main() {
	env := config.LoadEnv()

	router := config.InitializeServer()

	address := fmt.Sprintf("%s:%s", env.Host, env.Port)
	router.Run(address)
}
