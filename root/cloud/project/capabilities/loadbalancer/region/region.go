package region

import (
	"context"
	"fmt"

	"github.com/miton18/ovh-models/cloud"
	"github.com/miton18/ovh/metrics"
	ovhclient "github.com/ovh/go-ovh/ovh"
)

type Region interface {
	List(context.Context) ([]string, error)
	Get(context.Context, string) (*cloud.Region, error)
}

type region struct {
	client    *ovhclient.Client
	projectId string
}

func New(client *ovhclient.Client, project string) Region {
	return &region{client, project}
}

func (c *region) path() string {
	return fmt.Sprintf("/cloud/%s/capabilities/loadbalancer/region", c.projectId)
}

func (c *region) List(ctx context.Context) ([]string, error) {
	var regions []string

	oc := metrics.ObserveCall("ListLoadbalancerCapabilities")
	err := c.client.GetWithContext(ctx, c.path(), &regions)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return regions, nil
}

func (c *region) Get(ctx context.Context, regionName string) (*cloud.Region, error) {
	var region cloud.Region

	oc := metrics.ObserveCall("GetLoadbalancerRegion")
	err := c.client.GetWithContext(ctx, fmt.Sprintf(c.path()+"/"+regionName), &region)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return &region, nil
}
