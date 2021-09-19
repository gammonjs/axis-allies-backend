package logrus

import (
	"fmt"

	adaptee "github.com/sirupsen/logrus"
)

type Logger struct{}

func (logger Logger) Info(args ...interface{}) {
	fmt.Println("Here")
	adaptee.Info(args...)
}

func (logger Logger) Error(args ...interface{}) {
	adaptee.Error(args...)
}
