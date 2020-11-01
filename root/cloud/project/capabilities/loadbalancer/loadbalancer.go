package loadbalancer

import (
	"context"

	self "github.com/clever-telemetry/ovh/context"
	"github.com/clever-telemetry/ovh/root/cloud/project/capabilities/loadbalancer/region"
)

type Client interface {
	Region() region.Client
}

type client struct {
	ctx context.Context
}

func New(ctx context.Context) Client {
	ctx = self.AppendToPath(ctx, "/loadbalancer")

	return &client{
		ctx: ctx,
	}
}

func (c *client) Region() region.Client {
	return region.New(c.ctx)
}
