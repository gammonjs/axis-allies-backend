package gin

import adaptee "github.com/gin-gonic/gin"

func Create(...interface{}) (interface{}, error) {
	return &Router{Adaptee: adaptee.Default()}, nil
}
