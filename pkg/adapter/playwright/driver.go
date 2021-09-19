package playwright

import (
	"axis-allies-backend/pkg/contracts/web/browser"
	"errors"
	"fmt"

	adaptee "github.com/mxschmitt/playwright-go"
)

type Driver struct {
	Adaptee *adaptee.Playwright
}

func (self Driver) Browser() (browser.Instance, error) {
	options := adaptee.BrowserTypeLaunchOptions{
		Headless: adaptee.Bool(false),
	}

	browser, err := self.Adaptee.Chromium.Launch(options)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not launch browser: %v", err))
	}

	return &Browser{Adaptee: browser}, nil
}

func (self Driver) Stop() error {
	return self.Adaptee.Stop()
}
