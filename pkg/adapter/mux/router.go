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

// wrap a contract web context handler with a gin adapted handler
func (router Router) Middleware(handler func(api.Context)) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context := &Context{
				RequestAdaptee:        r,
				ResponseWriterAdaptee: w,
			}
			handler(context)
			next.ServeHTTP(w, r)
		})
	}
}

func (router Router) Handler(handler func(api.Context)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		adaptee := &Context{
			RequestAdaptee:        r,
			ResponseWriterAdaptee: w,
		}
		handler(adaptee)
	})
}

func (router Router) Use(middleware func(api.Context)) {
	router.Adaptee.Use(router.Middleware(middleware))
}

func (router Router) Get(relativePath string, handler func(api.Context)) {
	router.Adaptee.HandleFunc(relativePath, router.Handler(handler))
}
