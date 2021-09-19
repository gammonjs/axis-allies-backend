package browser

type Context interface {
	NewPage() (Page, error)
	Close() error
}
