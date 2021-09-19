package utility

type Configuration interface {
	Configure() error
	Get(key string) interface{}
	ServerUrl() string
}
