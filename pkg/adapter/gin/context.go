package gin

import (
	adaptee "github.com/gin-gonic/gin"
)

type Context struct {
	Adaptee *adaptee.Context
}
