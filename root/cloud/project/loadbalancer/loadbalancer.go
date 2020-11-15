package loadbalancer

import (
	"context"
	"fmt"
	"github.com/miton18/ovh/root/cloud/project/loadbalancer/configuration"

	"github.com/miton18/ovh-models/cloud"
	"github.com/miton18/ovh/metrics"
	ovhclient "github.com/ovh/go-ovh/ovh"

	"github.com/miton18/ovh/utils"
)

type Node interface {
	List(ctx context.Context) ([]string, error)
	// Fetch all LB
	Fetch(ctx context.Context) ([]cloud.LoadBalancer, error)
	Get(ctx context.Context, loadbalancerId string) (*cloud.LoadBalancer, error)
	// Create a new Node on the given region
	// name is optional
	// description is optional
	Create(ctx context.Context, region, name, description string) (*cloud.LoadBalancer, error)
	SetName(ctx context.Context, loadbalancerId, name string) (*cloud.LoadBalancer, error)
	SetDescription(ctx context.Context, loadbalancerId, description string) (*cloud.LoadBalancer, error)
	Delete(ctx context.Context, loadbalancerId string) error
	Configuration(loadbalancerId string) configuration.Node
}

type node struct {
	client    *ovhclient.Client
	project string
}

func New(client *ovhclient.Client, project string) Node {
	return &node{client, project}
}

func (l *node) path() string {
	return fmt.Sprintf("/cloud/project/%s/loadbalancer", l.project)
}

func (l *node) pathWithId(loadbalancerId string) string {
	return fmt.Sprintf("/cloud/project/%s/loadbalancer/%s", l.project, loadbalancerId)
}

func (l *node) List(ctx context.Context) ([]string, error) {
	var loadbalancers []string

	oc := metrics.ObserveCall("ListLoadbalancers")
	err := l.client.GetWithContext(ctx, l.path(), &loadbalancers)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return loadbalancers, nil
}

// WIP
func (l *node) Fetch(ctx context.Context) ([]cloud.LoadBalancer, error) {
	var loadbalancers []cloud.LoadBalancer

	l.client.Logger = utils.Logger{}
	req, err := l.client.NewRequest("GET", l.path(), nil, true)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Pagination-Mode", "CachedObjectList-Cursor")
	req.Header.Set("X-Pagination-Size", "10")

	res, err := l.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("bad stats %s", res.Status)
	}
	page := utils.NewPage(res.Header)
	fmt.Println(page.String())

	err = l.client.UnmarshalResponse(res, &loadbalancers)

	return loadbalancers, err
}

func (l *node) Get(ctx context.Context, loadbalancerId string) (*cloud.LoadBalancer, error) {
	var loadbalancer cloud.LoadBalancer

	oc := metrics.ObserveCall("GetLoadbalancer")
	err := l.client.GetWithContext(ctx, l.path()+"/"+loadbalancerId, &loadbalancer)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return &loadbalancer, nil
}

func (l *node) Create(ctx context.Context, region, name, description string) (*cloud.LoadBalancer, error) {
	create := cloud.LoadBalancerCreation{
		Region:      region,
		Name:        utils.NullableString(name),
		Description: utils.NullableString(description),
	}
	var loadbalancer cloud.LoadBalancer

	oc := metrics.ObserveCall("CreateLoadbalancer")
	err := l.client.PostWithContext(ctx, l.path(), &create, &loadbalancer)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return &loadbalancer, nil
}

func (l *node) SetName(ctx context.Context, loadbalancerId, name string) (*cloud.LoadBalancer, error) {
	update := utils.Map{
		"name": name,
	}
	var lb cloud.LoadBalancer

	oc := metrics.ObserveCall("SetLoadbalancerName")
	err := l.client.PutWithContext(ctx, l.pathWithId(loadbalancerId), &update, &lb)
	oc.End(err)
	if err != nil {
		return nil, err
	}
	return &lb, nil
}

func (l *node) SetDescription(ctx context.Context, loadbalancerId, description string) (*cloud.LoadBalancer, error) {
	update := utils.Map{
		"description": description,
	}
	var lb cloud.LoadBalancer

	oc := metrics.ObserveCall("SetLoadbalancerDescription")
	err := l.client.PutWithContext(ctx, l.pathWithId(loadbalancerId), &update, &lb)
	oc.End(err)
	if err != nil {
		return nil, err
	}
	return &lb, nil
}

func (l *node) Delete(ctx context.Context, loadbalancerId string) error {
	oc := metrics.ObserveCall("DeleteLoadbalancer")
	err := l.client.DeleteWithContext(ctx, l.path()+"/"+loadbalancerId, nil)
	oc.End(err)
	return err
}

func (l *node) Configuration(loadbalancerId string) configuration.Node {
	return configuration.New(l.client, l.project, loadbalancerId)
}
