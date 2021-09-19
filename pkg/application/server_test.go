package application

import (
	"axis-allies-backend/pkg/adapter/playwright"
	"testing"
	"time"
)

var Driver *playwright.DriverFactory = &playwright.DriverFactory{}

func TestHomePage(t *testing.T) {
	// arrange
	driver, _ := Driver.Create()
	browser, _ := driver.Browser()
	context, _ := browser.NewContext()
	page, _ := context.NewPage()

	defer context.Close()
	defer driver.Stop()

	// act
	page.Goto("localhost:8080")

	// assert
	time.Sleep(2 * time.Second)
}
