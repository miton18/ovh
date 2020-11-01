package cloud

import (
	"context"

	self "github.com/clever-telemetry/ovh/context"
	"github.com/clever-telemetry/ovh/root/cloud/project"
)

type Client interface {
	Project() project.Client
}

type client struct {
	ctx  context.Context
	path string
}

func New(ctx context.Context) Client {
	return &client{
		ctx:  ctx,
		path: "/cloud",
	}
}

func (c *client) Project() project.Client {
	ctx := self.WithPath(c.ctx, "/cloud")
	return project.New(ctx)
}
