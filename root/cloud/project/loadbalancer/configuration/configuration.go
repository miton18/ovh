package configuration

import (
	"context"
	"fmt"

	"github.com/miton18/ovh-models/cloud"
	"github.com/miton18/ovh/metrics"
	ovhclient "github.com/ovh/go-ovh/ovh"
)

type Node interface {
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

type node struct {
	client         *ovhclient.Client
	project      string
	loadbalancer string
}

func New(client *ovhclient.Client, project string, loadbalancer string) Node {
	return &node{client, project, loadbalancer}
}

func (c *node) path() string {
	return fmt.Sprintf("/cloud/project/%s/loadbalancer/%s/configuration", c.project, c.loadbalancer)
}

func (c *node) pathWithId(version int64) string {
	return fmt.Sprintf("/cloud/project/%s/loadbalancer/%s/configuration/%d", c.project, c.loadbalancer, version)
}

func (c *node) List(ctx context.Context) ([]int64, error) {
	var versions []int64

	oc := metrics.ObserveCall("ListLoadbalancerConfigs")
	err := c.client.GetWithContext(ctx, c.path(), &versions)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return versions, nil
}

func (c *node) Get(ctx context.Context, version int64) (*cloud.Configuration, error) {
	var config cloud.Configuration

	oc := metrics.ObserveCall("GetLoadbalancerConfig")
	err := c.client.GetWithContext(ctx, c.pathWithId(version), &config)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *node) Create(ctx context.Context, config *cloud.ConfigurationCreation) (*cloud.Configuration, error) {
	var _config cloud.Configuration

	oc := metrics.ObserveCall("CreateLoadbalancerConfig")
	err := c.client.PostWithContext(ctx, c.path(), config, &_config)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return &_config, nil
}

func (c *node) Delete(ctx context.Context, version int64) error {
	oc := metrics.ObserveCall("DeleteLoadbalancerConfig")
	err := c.client.DeleteWithContext(ctx, c.pathWithId(version), nil)
	oc.End(err)

	return err
}

func (c *node) Apply(ctx context.Context, version int64) error {
	oc := metrics.ObserveCall("ApplyLoadbalancerConfig")
	err := c.client.PostWithContext(ctx, c.pathWithId(version)+"/apply", nil, nil)
	oc.End(err)

	return err
}
