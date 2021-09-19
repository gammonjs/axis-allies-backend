package mux

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	RequestAdaptee        *http.Request
	ResponseWriterAdaptee http.ResponseWriter
}

func (context Context) Request() *http.Request {
	return context.RequestAdaptee
}

func (context Context) Response() http.ResponseWriter {
	return context.ResponseWriterAdaptee
}

func (context Context) JSON(status int, content interface{}) {
	json.NewEncoder(context.ResponseWriterAdaptee).Encode(content)
}
