package cloud

import (
	"context"

	"github.com/clever-telemetry/ovh/root/cloud/project"
)

type Client interface {
	Project() project.Client
}

type client struct {
	ctx context.Context
}

func New(ctx context.Context) Client {
	return &client{
		ctx: ctx,
	}
}

func (c *client) Project() project.Client {
	return project.New(c.ctx)
}
