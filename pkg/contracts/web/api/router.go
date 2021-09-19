package api

import "net/http"

type Router interface {
	http.Handler
	Use(middleware ...func(Context))
	Get(relativePath string, middleware ...func(Context))
}
