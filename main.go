package main

import (
	"fmt"
	"go-api/infra/config"
)

func main() {
	env := config.LoadEnv()

	router := config.InitializeServer()

	address := fmt.Sprintf("%s:%s", env.Host, env.Port)
	router.Run(address)
}
