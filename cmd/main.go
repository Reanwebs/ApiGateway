package main

import (
	"log"

	"gateway/pkg/auth/common/config"
	"gateway/pkg/auth/common/di"
)

func main() {
	c, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("cannot load config:", configerr)
	}

	server, dierr := di.InitializeAPI(c)
	if dierr != nil {
		log.Fatal("cannot initialize server", dierr)
	}
	server.Start()
}
