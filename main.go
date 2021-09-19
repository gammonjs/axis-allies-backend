package main

import (
	"axis-allies-backend/pkg/contracts/web/api"
	"axis-allies-backend/pkg/registry"

	"github.com/golobby/container/v2"
)

func main() {

	registry.Utility()
	registry.Router()
	registry.Server()

	var server api.Server
	container.Bind(&server)

	server.ListenAndServe()
}
