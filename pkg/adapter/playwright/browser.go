package playwright

import (
	"axis-allies-backend/pkg/contracts/web/browser"

	adaptee "github.com/mxschmitt/playwright-go"
)

type Browser struct {
	Adaptee adaptee.Browser
}

func (self Browser) NewContext() (browser.Context, error) {

	var options []adaptee.BrowserNewContextOptions
	browserContext, err := self.Adaptee.NewContext(options...)

	if err != nil {
		return nil, err
	}

	return &BrowserContext{Adaptee: browserContext}, nil
}
