package browser

type Page interface {
	Goto(url string) (Response, error)
	Text(text string) (interface{}, error)
}
