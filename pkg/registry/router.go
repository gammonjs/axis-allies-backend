package registry

import (
	"axis-allies-backend/pkg/adapter/gin"
	"axis-allies-backend/pkg/adapter/mux"
	"axis-allies-backend/pkg/contracts/web/api"

	"github.com/golobby/container/v2"
)

func Router() {
	container.Transient(func() api.Router {
		return Mux()
	})
}

func Gin() api.Router {
	construct, _ := gin.Create()
	router := construct.(*gin.Router)
	container.Bind(&router.Log)
	return router
}

func Mux() api.Router {
	construct, _ := mux.Create()
	router := construct.(*mux.Router)
	container.Bind(&router.Log)
	return router
}
