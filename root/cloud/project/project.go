package project

import (
	"context"

	"github.com/clever-telemetry/ovh-models/cloud"
	self "github.com/clever-telemetry/ovh/context"
	"github.com/clever-telemetry/ovh/root/cloud/project/capabilities"
	"github.com/clever-telemetry/ovh/root/cloud/project/loadbalancer"
)

type Client interface {
	List(context.Context) ([]string, error)
	Get(context.Context, string) (*cloud.Project, error)
	Loadbalancer(string) loadbalancer.Client
	Capabilities(string) capabilities.Client
}

type client struct {
	ctx  context.Context
	path string
}

func New(ctx context.Context) Client {
	ctx = self.AppendToPath(ctx, "/project")

	return &client{
		ctx:  ctx,
		path: self.Path(ctx),
	}
}

func (c *client) List(ctx context.Context) ([]string, error) {
	var projects []string

	if err := self.Client(c.ctx).GetWithContext(ctx, c.path, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (c *client) Get(ctx context.Context, projectId string) (*cloud.Project, error) {
	var project cloud.Project

	if err := self.Client(c.ctx).GetWithContext(ctx, c.path+"/"+projectId, &project); err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *client) Loadbalancer(project string) loadbalancer.Client {
	ctx := self.WithProjectId(c.ctx, project)
	ctx = self.AppendToPath(ctx, "/"+project)
	return loadbalancer.New(ctx)
}

func (c *client) Capabilities(project string) capabilities.Client {
	ctx := self.WithProjectId(c.ctx, project)
	ctx = self.AppendToPath(ctx, "/"+project)
	return capabilities.New(ctx)
}
