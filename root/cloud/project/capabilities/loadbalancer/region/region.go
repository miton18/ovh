package region

import (
	"context"
	"fmt"

	"github.com/clever-telemetry/ovh-models/cloud"
	self "github.com/clever-telemetry/ovh/context"
	"github.com/clever-telemetry/ovh/metrics"
	ovhclient "github.com/ovh/go-ovh/ovh"
)

type Client interface {
	List(context.Context) ([]string, error)
	Get(context.Context, string) (*cloud.Region, error)
}

type client struct {
	ctx       context.Context
	client    *ovhclient.Client
	projectId string
}

func New(ctx context.Context) Client {
	return &client{
		ctx:       ctx,
		client:    self.Client(ctx),
		projectId: self.ProjectId(ctx),
	}
}

func (c *client) path() string {
	return fmt.Sprintf("/cloud/%s/capabilities/loadbalancer/region", c.projectId)
}

func (c *client) List(ctx context.Context) ([]string, error) {
	var regions []string

	oc := metrics.ObserveCall("ListLoadbalancerCapabilities")
	err := c.client.GetWithContext(ctx, c.path(), &regions)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return regions, nil
}

func (c *client) Get(ctx context.Context, regionName string) (*cloud.Region, error) {
	var region cloud.Region

	oc := metrics.ObserveCall("GetLoadbalancerRegion")
	err := c.client.GetWithContext(ctx, fmt.Sprintf(c.path()+"/"+regionName), &region)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return &region, nil
}
