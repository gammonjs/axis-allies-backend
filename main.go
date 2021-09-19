package main

import (
	"axis-allies-backend/pkg/contracts/utility"
	"axis-allies-backend/pkg/contracts/web/api"
	"axis-allies-backend/pkg/registry"

	"github.com/golobby/container/v2"
)

func main() {

	registry.Utility()
	registry.Router()
	registry.Server()

	var configuration utility.Configuration
	container.Bind(&configuration)

	var log utility.Logger
	container.Bind(&log)

	if err := configuration.Configure(); err != nil {
		log.Error(err.Error())
		return
	}

	var server api.Server
	container.Bind(&server)

	server.ListenAndServe()
}
