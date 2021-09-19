package registry

import (
	"axis-allies-backend/pkg/adapter/logrus"
	"axis-allies-backend/pkg/contracts/utility"

	"github.com/golobby/container/v2"
)

func Utility() {
	container.Singleton(func() utility.Logger {
		return &logrus.Logger{}
	})
}
