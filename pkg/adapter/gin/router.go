package gin

import (
	"axis-allies-backend/pkg/contracts/utility"
	"axis-allies-backend/pkg/contracts/web/api"
	"net/http"

	adaptee "github.com/gin-gonic/gin"
)

type Router struct {
	Adaptee *adaptee.Engine
	Log     utility.Logger
}

func (router Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	router.Adaptee.ServeHTTP(response, request)
}

// wrap a contract web context handler with a gin adapted handler
func (router Router) Handler(handler func(api.Context)) adaptee.HandlerFunc {
	return func(adapteeContext *adaptee.Context) {
		adaptee := &Context{Adaptee: adapteeContext}
		handler(adaptee)
	}
}

func (router Router) Use(middleware func(api.Context)) {
	router.Adaptee.Use(router.Handler(middleware))
}

func (router Router) Get(relativePath string, handler func(api.Context)) {
	router.Adaptee.GET(relativePath, router.Handler(handler))
}
