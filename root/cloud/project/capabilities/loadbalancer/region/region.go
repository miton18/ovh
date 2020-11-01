package region

import (
	"context"
	"fmt"

	"github.com/clever-telemetry/ovh-models/cloud"
	self "github.com/clever-telemetry/ovh/context"
	ovhclient "github.com/ovh/go-ovh/ovh"
)

type Client interface {
	List(context.Context) ([]string, error)
	Get(context.Context, string) (*cloud.Region, error)
}

type client struct {
	ctx    context.Context
	client *ovhclient.Client
	path   string
}

func New(ctx context.Context) Client {
	ctx = self.AppendToPath(ctx, "/region")
	return &client{
		ctx:    ctx,
		client: self.Client(ctx),
		path:   self.Path(ctx),
	}
}

func (c *client) List(ctx context.Context) ([]string, error) {
	var regions []string

	if err := c.client.GetWithContext(ctx, c.path, &regions); err != nil {
		return nil, err
	}

	return regions, nil
}

func (c *client) Get(ctx context.Context, regionName string) (*cloud.Region, error) {
	var region cloud.Region

	if err := c.client.GetWithContext(ctx, fmt.Sprintf(c.path+"/"+regionName), &region); err != nil {
		return nil, err
	}

	return &region, nil
}
