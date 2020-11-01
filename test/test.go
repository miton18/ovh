package test

import (
	"os"

	ovhclient "github.com/ovh/go-ovh/ovh"
)

var Client *ovhclient.Client

func init() {
	var err error
	Client, err = ovhclient.NewClient(
		os.Getenv("OVH_ENDPOINT"),
		os.Getenv("OVH_APP_KEY"),
		os.Getenv("OVH_APP_SECRET"),
		os.Getenv("OVH_CONSUMER_KEY"),
	)

	if err != nil {
		panic(err.Error())
	}
}
