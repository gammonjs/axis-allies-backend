package viper

import (
	"fmt"

	adaptee "github.com/spf13/viper"
)

const ApplicationPath string = "./env"
const IntegrationPath string = "../env"

type Configuration struct{}

func (configuration Configuration) Configure() error {
	adaptee.SetConfigName("local")
	adaptee.SetConfigType("yaml")
	adaptee.AddConfigPath(ApplicationPath)
	adaptee.AddConfigPath(IntegrationPath)

	err := adaptee.ReadInConfig() // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		return err
	}

	return nil
}

func (configuration Configuration) Get(key string) interface{} {
	return adaptee.Get(key)
}

func (configuration Configuration) ServerUrl() string {
	server := adaptee.Get("server").(map[string]interface{})
	return fmt.Sprintf("%s:%s", server["address"], server["port"])
}
