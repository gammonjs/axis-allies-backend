package gin

import (
	"net/http"

	adaptee "github.com/gin-gonic/gin"
)

type Context struct {
	Adaptee *adaptee.Context
}

func (context Context) Request() *http.Request {
	return context.Adaptee.Request
}

func (context Context) Response() http.ResponseWriter {
	return context.Adaptee.Writer
}

func (context Context) JSON(status int, content interface{}) {
	context.Adaptee.JSON(status, content)
}
