package mux

import (
	adaptee "github.com/gorilla/mux"
)

func Create(...interface{}) (interface{}, error) {
	server := &Router{Adaptee: adaptee.NewRouter()}
	return server, nil
}
