package capabilities

import (
	"context"

	"github.com/clever-telemetry/ovh/root/cloud/project/capabilities/loadbalancer"
)

type Client interface {
	Loadbalancer() loadbalancer.Client
}

type client struct {
	ctx context.Context
}

func New(ctx context.Context) Client {
	return &client{ctx}
}

func (c *client) Loadbalancer() loadbalancer.Client {
	return loadbalancer.New(c.ctx)
}
