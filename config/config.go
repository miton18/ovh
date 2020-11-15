package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Configuration struct {
	Endpoint          Endpoint `json:"endpoint"`
	ApplicationKey    string   `json:"applicationKey"`
	ApplicationSecret string   `json:"applicationSecret"`
	ConsumerKey       string   `json:"consumerKey"`
}

type loader func() (*Configuration, error)

func Must(l loader) *Configuration {
	c, err := l()
	if err != nil {
		panic(err.Error())
	}

	return c
}

// Read configuration from env vars
func FromEnv() (*Configuration, error) {
	return &Configuration{
		Endpoint: Endpoint(os.Getenv("OVH_ENDPOINT")),
		ApplicationKey: os.Getenv("OVH_APPLICATION_KEY"),
		ApplicationSecret: os.Getenv("OVH_APPLICATION_SECRET"),
		ConsumerKey: os.Getenv("OVH_CONSUMER_KEY"),
	}, nil
}

// Read configuration from predefined or provided paths
func FromFile(path string) (*Configuration, error) {
	var c Configuration
	v := viper.New()
	v.SetConfigName("ovh")
	v.AddConfigPath(path)
	v.AddConfigPath(".")
	v.AddConfigPath("$HOME/.config/ovh")
	v.AddConfigPath("$HOME/.ovh")
	v.AddConfigPath("$HOME")
	v.AddConfigPath("/etc/ovh")
	v.AddConfigPath("/etc")
	v.SetDefault("endpoint", OVH_EU)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

func Auto() (*Configuration, error) {
	c, err := FromFile("")
	if err == nil {
		return c, nil
	}
	fmt.Printf("%+v\n", err.Error())

	c, err = FromEnv()
	if err == nil {
		return c, nil
	}

	return nil, err
}
