package playwright

import (
	"axis-allies-backend/pkg/contracts/web/browser"

	adaptee "github.com/mxschmitt/playwright-go"
)

type DriverFactory struct{}

func (self DriverFactory) Create(objects ...interface{}) (browser.Driver, error) {
	adaptee, err := adaptee.Run()

	if err != nil {
		return nil, err
	}

	return &Driver{Adaptee: adaptee}, nil
}
