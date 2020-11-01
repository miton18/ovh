package loadbalancer

import (
	"context"

	"github.com/clever-telemetry/ovh-models/cloud"
	self "github.com/clever-telemetry/ovh/context"
	ovhclient "github.com/ovh/go-ovh/ovh"

	"github.com/clever-telemetry/ovh/utils"
)

type Client interface {
	List(context.Context) ([]string, error)
	// Fetch all LB
	Fetch(ctx context.Context) ([]cloud.LoadBalancer, error)
	Get(context.Context, string) (*cloud.LoadBalancer, error)
	// Create a new Loadbalancer on the given region
	// name is optional
	// description is optional
	Create(context.Context, string, string, string) (*cloud.LoadBalancer, error)
}

type client struct {
	ctx    context.Context
	client *ovhclient.Client
	path   string
}

func New(ctx context.Context) Client {
	ctx = self.AppendToPath(ctx, "/loadbalancer")

	return &client{
		ctx:    ctx,
		client: self.Client(ctx),
		path:   self.Path(ctx),
	}
}

func (c *client) List(ctx context.Context) ([]string, error) {
	var loadbalancers []string

	if err := c.client.GetWithContext(ctx, c.path, &loadbalancers); err != nil {
		return nil, err
	}

	return loadbalancers, nil
}

func (c *client) Fetch(ctx context.Context) ([]cloud.LoadBalancer, error) {
	var loadbalancers []cloud.LoadBalancer

	req, err := c.client.NewRequest("GET", c.path, nil, true)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Pagination-Mode", "CachedObjectList-Pages")
	req.Header.Set("X-Pagination-Size", "10000")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = c.client.UnmarshalResponse(res, &loadbalancers)

	return loadbalancers, err
}

func (c *client) Get(ctx context.Context, loadbalancerId string) (*cloud.LoadBalancer, error) {
	var loadbalancer cloud.LoadBalancer

	if err := c.client.GetWithContext(ctx, c.path+"/"+loadbalancerId, &loadbalancer); err != nil {
		return nil, err
	}

	return &loadbalancer, nil
}

func (c *client) Create(ctx context.Context, region, name, description string) (*cloud.LoadBalancer, error) {
	create := cloud.LoadBalancerCreation{
		Region:      region,
		Name:        utils.NullableString(name),
		Description: utils.NullableString(description),
	}
	var loadbalancer cloud.LoadBalancer

	if err := c.client.PostWithContext(ctx, c.path, &create, &loadbalancer); err != nil {
		return nil, err
	}

	return &loadbalancer, nil
}
