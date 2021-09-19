package playwright

import (
	"axis-allies-backend/pkg/contracts/web/browser"

	adaptee "github.com/mxschmitt/playwright-go"
)

type Page struct {
	Adaptee adaptee.Page
}

func (self Page) Goto(url string) (browser.Response, error) {
	var options []adaptee.PageGotoOptions
	response, err := self.Adaptee.Goto(url, options...)

	if err != nil {
		return nil, err
	}

	return &Response{Adaptee: response}, nil
}

func (page Page) EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error) {
	return page.Adaptee.EvalOnSelectorAll(selector, expression, options...)
}
