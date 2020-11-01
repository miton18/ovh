package capabilities

import (
	"context"

	self "github.com/clever-telemetry/ovh/context"
	"github.com/clever-telemetry/ovh/root/cloud/project/capabilities/loadbalancer"
)

type Client interface {
	Loadbalancer() loadbalancer.Client
}

type client struct {
	ctx context.Context
}

func New(ctx context.Context) Client {
	ctx = self.AppendToPath(ctx, "/capabilities")

	return &client{
		ctx: ctx,
	}
}

func (c *client) Loadbalancer() loadbalancer.Client {
	return loadbalancer.New(c.ctx)
}
