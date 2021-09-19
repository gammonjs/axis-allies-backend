package browser

type Page interface {
	Goto(url string) (Response, error)
	EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
}
