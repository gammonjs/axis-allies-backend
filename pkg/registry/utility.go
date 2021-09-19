package registry

import (
	"axis-allies-backend/pkg/adapter/logrus"
	"axis-allies-backend/pkg/adapter/viper"
	"axis-allies-backend/pkg/contracts/utility"

	"github.com/golobby/container/v2"
)

func Utility() {
	container.Singleton(func() utility.Configuration {
		return &viper.Configuration{}
	})

	container.Singleton(func() utility.Logger {
		return &logrus.Logger{}
	})

	var configuration utility.Configuration
	container.Bind(&configuration)

	var log utility.Logger
	container.Bind(&log)

	if err := configuration.Configure(); err != nil {
		log.Error(err.Error())
	}
}
