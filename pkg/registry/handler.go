package registry

import (
	"axis-allies-backend/pkg/contracts/web/api"
	"axis-allies-backend/pkg/handler"

	"github.com/golobby/container/v2"
)

func Handler() {
	container.Transient(func() api.HomeHandler {
		return &handler.Home{}
	})
}
