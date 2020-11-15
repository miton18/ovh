package loadbalancer

import (
	"context"
	"fmt"

	"github.com/clever-telemetry/ovh-models/cloud"
	self "github.com/clever-telemetry/ovh/context"
	"github.com/clever-telemetry/ovh/metrics"
	ovhclient "github.com/ovh/go-ovh/ovh"

	"github.com/clever-telemetry/ovh/utils"
)

type Client interface {
	List(ctx context.Context) ([]string, error)
	// Fetch all LB
	Fetch(ctx context.Context) ([]cloud.LoadBalancer, error)
	Get(ctx context.Context, loadbalancerId string) (*cloud.LoadBalancer, error)
	// Create a new Loadbalancer on the given region
	// name is optional
	// description is optional
	Create(ctx context.Context, region, name, description string) (*cloud.LoadBalancer, error)
	SetName(ctx context.Context, loadbalancerId, name string) (*cloud.LoadBalancer, error)
	SetDescription(ctx context.Context, loadbalancerId, description string) (*cloud.LoadBalancer, error)
	Delete(ctx context.Context, loadbalancerId string) error
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
	return fmt.Sprintf("/cloud/project/%s/loadbalancer", c.projectId)
}

func (c *client) pathWithId(loadbalancerId string) string {
	return fmt.Sprintf("/cloud/project/%s/loadbalancer/%s", c.projectId, loadbalancerId)
}

func (c *client) List(ctx context.Context) ([]string, error) {
	var loadbalancers []string

	oc := metrics.ObserveCall("ListLoadbalancers")
	err := c.client.GetWithContext(ctx, c.path(), &loadbalancers)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return loadbalancers, nil
}

// WIP
func (c *client) Fetch(ctx context.Context) ([]cloud.LoadBalancer, error) {
	var loadbalancers []cloud.LoadBalancer

	c.client.Logger = utils.Logger{}
	req, err := c.client.NewRequest("GET", c.path(), nil, true)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Pagination-Mode", "CachedObjectList-Cursor")
	req.Header.Set("X-Pagination-Size", "10")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("bad stats %s", res.Status)
	}
	page := utils.NewPage(res.Header)
	fmt.Println(page.String())

	err = c.client.UnmarshalResponse(res, &loadbalancers)

	return loadbalancers, err
}

func (c *client) Get(ctx context.Context, loadbalancerId string) (*cloud.LoadBalancer, error) {
	var loadbalancer cloud.LoadBalancer

	oc := metrics.ObserveCall("GetLoadbalancer")
	err := c.client.GetWithContext(ctx, c.path()+"/"+loadbalancerId, &loadbalancer)
	oc.End(err)
	if err != nil {
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

	oc := metrics.ObserveCall("CreateLoadbalancer")
	err := c.client.PostWithContext(ctx, c.path(), &create, &loadbalancer)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return &loadbalancer, nil
}

func (c *client) SetName(ctx context.Context, loadbalancerId, name string) (*cloud.LoadBalancer, error) {
	update := utils.Map{
		"name": name,
	}
	var lb cloud.LoadBalancer

	oc := metrics.ObserveCall("SetLoadbalancerName")
	err := c.client.PutWithContext(ctx, c.pathWithId(loadbalancerId), &update, &lb)
	oc.End(err)
	if err != nil {
		return nil, err
	}
	return &lb, nil
}

func (c *client) SetDescription(ctx context.Context, loadbalancerId, description string) (*cloud.LoadBalancer, error) {
	update := utils.Map{
		"description": description,
	}
	var lb cloud.LoadBalancer

	oc := metrics.ObserveCall("SetLoadbalancerDescription")
	err := c.client.PutWithContext(ctx, c.pathWithId(loadbalancerId), &update, &lb)
	oc.End(err)
	if err != nil {
		return nil, err
	}
	return &lb, nil
}

func (c *client) Delete(ctx context.Context, loadbalancerId string) error {
	oc := metrics.ObserveCall("DeleteLoadbalancer")
	err := c.client.DeleteWithContext(ctx, c.path()+"/"+loadbalancerId, nil)
	oc.End(err)
	return err
}
