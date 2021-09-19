package integration

import (
	"axis-allies-backend/pkg/adapter/playwright"
	"axis-allies-backend/pkg/contracts/web/browser"
)

type Fixture struct {
	Driver   browser.Driver
	Instance browser.Instance
	Context  browser.Context
}

type FixtureFactory struct{}

func (factory FixtureFactory) Create() *Fixture {
	driverFactory := &playwright.DriverFactory{}
	driver, _ := driverFactory.Create()
	browser, _ := driver.Browser()
	context, _ := browser.NewContext()

	return &Fixture{
		Driver:   driver,
		Instance: browser,
		Context:  context,
	}
}

func Cleanup(fixture *Fixture) {
	fixture.Context.Close()
	fixture.Driver.Stop()
}
