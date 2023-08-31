package main

import (
	"log"

	"gateway/docs"
	"gateway/pkg/common/config"
	"gateway/pkg/common/di"
)

//	@title			REAN WEB CONNECT
//	@version		2.0
//	@description	REAN WEB CONNECT FOR VIDEO CONFERENCING AND BROADCAST.

//	@contact
// name: REAN WEB CONNECT
// url: https://github.com/Reanwebs
// email: reanwebpvt@gmail.com

//	@license
// name: nil
// url: nil

//	@host	nil

// @Basepath	/
// @Accept		json
// @Produce	json
// @Router		/ [get]
func main() {

	c, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("cannot load config:", configerr)
	}

	docs.SwaggerInfo.Title = c.SwagTitle
	docs.SwaggerInfo.Description = c.SwagDescription
	docs.SwaggerInfo.Version = c.SwagVersion

	server, dierr := di.InitializeAPI(c)
	if dierr != nil {
		log.Fatal("cannot initialize server", dierr)
	}
	server.Start()
}
