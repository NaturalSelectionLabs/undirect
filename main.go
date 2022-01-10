package main

import (
	"github.com/Candinya/undirect/config"
	"github.com/Candinya/undirect/routers"
	"github.com/Candinya/undirect/routers/url"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	// Load config

	log.Println("Loading config...")

	if err := config.LoadConfig("config.yml"); err != nil {
		panic(err)
	}

	log.Println("Config loaded successfully.")

	if config.GlobalConfig.Debug {
		log.Println("Working on debug mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init routers

	log.Println("Initializing routers...")

	routers.Include(url.Routers)

	r := routers.Init()

	log.Println("Routers initialized successfully.")

	log.Println("Starting gin server...")

	if err := r.Run(config.GlobalConfig.Listen); err != nil {
		panic(err)
	}

}
