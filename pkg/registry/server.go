package registry

import (
	"axis-allies-backend/pkg/application"
	"axis-allies-backend/pkg/contracts/web/api"

	"github.com/golobby/container/v2"
)

func Server() {
	container.Transient(func() api.Server {
		server := &application.Server{}
		container.Bind(&server.Router)
		container.Bind(&server.Log)
		return server
	})
}
