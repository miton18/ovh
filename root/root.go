package root

import (
	"context"

	"github.com/clever-telemetry/ovh/root/cloud"
)

type Client interface {
	Cloud() cloud.Client
}

type client struct {
	ctx context.Context
}

func New(ctx context.Context) Client {
	return &client{ctx}
}

func (c *client) Cloud() cloud.Client {
	return cloud.New(c.ctx)
}
