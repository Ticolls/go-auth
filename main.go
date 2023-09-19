package main

import (
	"github.com/Ticolls/go-auth/config"
	"github.com/Ticolls/go-auth/router"
)

func main() {

	logger := *config.GetLogger("main")

	// Initialize config
	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization erro: %v", err)
		return
	}

	router.Init()
}
