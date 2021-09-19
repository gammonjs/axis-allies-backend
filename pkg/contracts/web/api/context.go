package api

import "net/http"

type Context interface {
	Request() *http.Request
	Response() http.ResponseWriter
	JSON(status int, content interface{})
}
