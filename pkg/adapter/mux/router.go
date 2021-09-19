package mux

import (
	"axis-allies-backend/pkg/contracts/utility"
	"axis-allies-backend/pkg/contracts/web/api"
	"net/http"

	adaptee "github.com/gorilla/mux"
)

type Router struct {
	Adaptee *adaptee.Router
	Log     utility.Logger
}

func (router Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	router.Adaptee.ServeHTTP(response, request)
}

func (router Router) Use(middleware ...func(api.Context)) {

}

func (router Router) Get(relativePath string, middleware ...func(api.Context)) {

}
