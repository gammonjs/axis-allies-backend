package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testFixture *FixtureFactory = &FixtureFactory{}

func TestHomePage(t *testing.T) {
	// setup and cleanup
	fixture := testFixture.Create()
	defer Cleanup(fixture)

	// arrange
	page, _ := fixture.Context.NewPage()

	// act
	page.Goto("localhost:8080")
	count, err := page.EvalOnSelectorAll("text={\"Hello\":\"World\"}", "el => el.length")

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, count, 1)
}
