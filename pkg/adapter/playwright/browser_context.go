package playwright

import (
	"axis-allies-backend/pkg/contracts/web/browser"

	adaptee "github.com/mxschmitt/playwright-go"
)

type BrowserContext struct {
	Adaptee adaptee.BrowserContext
}

func (self BrowserContext) NewPage() (browser.Page, error) {
	page, err := self.Adaptee.NewPage()

	if err != nil {
		return nil, err
	}

	return &Page{Adaptee: page}, nil
}

func (self BrowserContext) Close() error {
	return self.Adaptee.Close()
}
