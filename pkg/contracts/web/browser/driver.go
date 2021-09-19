package browser

type Driver interface {
	Browser() (Instance, error)
	Stop() error
}
