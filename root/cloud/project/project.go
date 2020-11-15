package project

import (
	"context"

	"github.com/clever-telemetry/ovh-models/cloud"
	self "github.com/clever-telemetry/ovh/context"
	"github.com/clever-telemetry/ovh/metrics"
	"github.com/clever-telemetry/ovh/root/cloud/project/capabilities"
	"github.com/clever-telemetry/ovh/root/cloud/project/loadbalancer"
	ovhclient "github.com/ovh/go-ovh/ovh"
)

type Client interface {
	List(context.Context) ([]string, error)
	Get(context.Context, string) (*cloud.Project, error)
	Loadbalancer(string) loadbalancer.Client
	Capabilities(string) capabilities.Client
}

type client struct {
	ctx    context.Context
	client *ovhclient.Client
}

func New(ctx context.Context) Client {
	return &client{
		ctx:    ctx,
		client: self.Client(ctx),
	}
}

func (c *client) List(ctx context.Context) ([]string, error) {
	var projects []string

	oc := metrics.ObserveCall("ListProjects")
	err := c.client.GetWithContext(ctx, "/cloud/project", &projects)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (c *client) Get(ctx context.Context, projectId string) (*cloud.Project, error) {
	var project cloud.Project

	oc := metrics.ObserveCall("GetProject")
	err := c.client.GetWithContext(ctx, "/cloud/project/"+projectId, &project)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *client) Loadbalancer(project string) loadbalancer.Client {
	ctx := self.WithProjectId(c.ctx, project)
	return loadbalancer.New(ctx)
}

func (c *client) Capabilities(project string) capabilities.Client {
	ctx := self.WithProjectId(c.ctx, project)
	return capabilities.New(ctx)
}
