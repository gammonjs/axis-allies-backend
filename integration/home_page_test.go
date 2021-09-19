package integration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	Runner(func() {
		// arrange
		render := "{\"Hello\":\"World\"}"
		expected := 1
		page, _ := fixture.Context.NewPage()

		// act
		page.Goto(fixture.Configuration.ServerUrl())
		actual, err := page.Text(render)

		// assert
		assert.Equal(t, err, nil)
		assert.Equal(t, actual, expected, fmt.Sprintf("expected the following page render: %s", render))
	})
}
