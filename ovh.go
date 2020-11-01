package ovh

import (
	"context"
	"os"

	ovhclient "github.com/ovh/go-ovh/ovh"

	selfcontext "github.com/clever-telemetry/ovh/context"
	"github.com/clever-telemetry/ovh/root"
	"github.com/clever-telemetry/ovh/utils"
)

var (
	String = utils.NullableString
)

func New(endpoint Endpoint, appKey, appSecret, consumerKey string) (root.Client, error) {

	_client, err := ovhclient.NewClient(
		endpoint.String(),
		appKey,
		appSecret,
		consumerKey,
	)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	ctx = selfcontext.WithClient(ctx, _client)
	return root.New(ctx), nil
}

func MustNew(endpoint Endpoint, appKey, appSecret, consumerKey string) root.Client {
	c, err := New(endpoint, appKey, appSecret, consumerKey)
	if err != nil {
		panic(err.Error())
	}

	return c
}

func NewFromEnv() (root.Client, error) {
	return New(
		Endpoint(os.Getenv("OVH_ENDPOINT")),
		os.Getenv("OVH_APP_KEY"),
		os.Getenv("OVH_APP_SECRET"),
		os.Getenv("OVH_CONSUMER_KEY"),
	)
}

func MustNewFromEnv() root.Client {
	c, err := NewFromEnv()
	if err != nil {
		panic(err.Error())
	}

	return c
}
