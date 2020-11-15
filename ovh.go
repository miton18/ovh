package ovh

import (
	"github.com/miton18/ovh/config"
	ovhclient "github.com/ovh/go-ovh/ovh"

	"github.com/miton18/ovh/root"
	"github.com/miton18/ovh/utils"
)

var (
	String = utils.NullableString
)

type SDK root.Node

func New(config *config.Configuration) (SDK, error) {

	_client, err := ovhclient.NewClient(
		config.Endpoint.String(),
		config.ApplicationKey,
		config.ApplicationSecret,
		config.ConsumerKey,
	)
	if err != nil {
		return nil, err
	}

	return root.New(_client), nil
}
