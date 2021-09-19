package utility

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}
