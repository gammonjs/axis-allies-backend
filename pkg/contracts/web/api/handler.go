package api

type Handler interface {
	Handler(Context)
}

type HomeHandler interface {
	Handler
}
