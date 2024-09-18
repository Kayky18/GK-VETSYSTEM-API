package main

import (
	"github.com/Kayky18/GK_API/config"
	"github.com/Kayky18/GK_API/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	//Initialize DB and Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("ERROR TO INITIALIZE INIT FUNC: %v", err)
		return
	}

	//Initialize Router
	router.InitializeRouter()
}
