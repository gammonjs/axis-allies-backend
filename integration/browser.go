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

	fixture := &Fixture{
		Driver:   driver,
		Instance: browser,
		Context:  context,
	}

	container.Bind(&fixture.Configuration)
	container.Bind(&fixture.Log)

	return fixture
}

func Cleanup(fixture *Fixture) {
	fixture.Context.Close()
	fixture.Driver.Stop()
}
