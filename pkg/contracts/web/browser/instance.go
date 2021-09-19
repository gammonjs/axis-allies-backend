package browser

type Instance interface {
	NewContext() (Context, error)
}
