package configuration

import (
	"context"
	"fmt"

	"github.com/clever-telemetry/ovh-models/cloud"
	self "github.com/clever-telemetry/ovh/context"
	"github.com/clever-telemetry/ovh/metrics"
	ovhclient "github.com/ovh/go-ovh/ovh"
)

type Client interface {
	List(ctx context.Context) ([]int64, error)
	// Fetch all LB
	Get(ctx context.Context, version int64) (*cloud.Configuration, error)
	// Create a new Loadbalancer on the given region
	// name is optional
	// description is optional
	Create(ctx context.Context, config *cloud.ConfigurationCreation) (*cloud.Configuration, error)
	Delete(ctx context.Context, version int64) error
	Apply(ctx context.Context, version int64) error
}

type client struct {
	ctx            context.Context
	client         *ovhclient.Client
	projectId      string
	loadbalancerId string
}

func New(ctx context.Context) Client {
	return &client{
		ctx:            ctx,
		client:         self.Client(ctx),
		projectId:      self.ProjectId(ctx),
		loadbalancerId: self.LoadbalancerId(ctx),
	}
}

func (c *client) path() string {
	return fmt.Sprintf("/cloud/project/%s/loadbalancer/%s/configuration", c.projectId, c.loadbalancerId)
}

func (c *client) pathWithId(version int64) string {
	return fmt.Sprintf("/cloud/project/%s/loadbalancer/%s/configuration/%d", c.projectId, c.loadbalancerId, version)
}

func (c *client) List(ctx context.Context) ([]int64, error) {
	var versions []int64

	oc := metrics.ObserveCall("ListLoadbalancerConfigs")
	err := c.client.GetWithContext(ctx, c.path(), &versions)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return versions, nil
}

func (c *client) Get(ctx context.Context, version int64) (*cloud.Configuration, error) {
	var config cloud.Configuration

	oc := metrics.ObserveCall("GetLoadbalancerConfig")
	err := c.client.GetWithContext(ctx, c.pathWithId(version), &config)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *client) Create(ctx context.Context, config *cloud.ConfigurationCreation) (*cloud.Configuration, error) {
	var _config cloud.Configuration

	oc := metrics.ObserveCall("CreateLoadbalancerConfig")
	err := c.client.PostWithContext(ctx, c.path(), config, &_config)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return &_config, nil
}

func (c *client) Delete(ctx context.Context, version int64) error {
	oc := metrics.ObserveCall("DeleteLoadbalancerConfig")
	err := c.client.DeleteWithContext(ctx, c.pathWithId(version), nil)
	oc.End(err)

	return err
}

func (c *client) Apply(ctx context.Context, version int64) error {
	oc := metrics.ObserveCall("ApplyLoadbalancerConfig")
	err := c.client.PostWithContext(ctx, c.pathWithId(version)+"/apply", nil, nil)
	oc.End(err)

	return err
}
