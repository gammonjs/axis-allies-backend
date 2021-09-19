package integration

import (
	"axis-allies-backend/pkg/adapter/playwright"
	"axis-allies-backend/pkg/contracts/utility"
	"axis-allies-backend/pkg/contracts/web/browser"
	"axis-allies-backend/pkg/registry"

	"github.com/golobby/container/v2"
)

type Fixture struct {
	Configuration utility.Configuration
	Log           utility.Logger
	Driver        browser.Driver
	Instance      browser.Instance
	Context       browser.Context
}

type FixtureFactory struct{}

func (factory FixtureFactory) Create() *Fixture {
	driverFactory := &playwright.DriverFactory{}
	driver, _ := driverFactory.Create()
	browser, _ := driver.Browser()
	context, _ := browser.NewContext()

	registry.Utility()

	var configuration utility.Configuration
	container.Bind(&configuration)

	var log utility.Logger
	container.Bind(&log)

	if err := configuration.Configure(); err != nil {
		log.Error(err.Error())
		return nil
	}

	return &Fixture{
		Configuration: configuration,
		Log:           log,
		Driver:        driver,
		Instance:      browser,
		Context:       context,
	}
}

func Cleanup(fixture *Fixture) {
	fixture.Context.Close()
	fixture.Driver.Stop()
}
