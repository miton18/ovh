package project

import (
	"context"

	"github.com/miton18/ovh-models/cloud"
	"github.com/miton18/ovh/metrics"
	"github.com/miton18/ovh/root/cloud/project/capabilities"
	"github.com/miton18/ovh/root/cloud/project/loadbalancer"
	ovhclient "github.com/ovh/go-ovh/ovh"
)

type Node interface {
	List(context.Context) ([]string, error)
	Get(context.Context, string) (*cloud.Project, error)
	Loadbalancer(string) loadbalancer.Node
	Capabilities(string) capabilities.Capabilities
}

type node struct {
	client *ovhclient.Client
}

func New(client *ovhclient.Client) Node {
	return &node{client}
}

func (p *node) List(ctx context.Context) ([]string, error) {
	var projects []string

	oc := metrics.ObserveCall("ListProjects")
	err := p.client.GetWithContext(ctx, "/cloud/project", &projects)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (p *node) Get(ctx context.Context, projectId string) (*cloud.Project, error) {
	var project cloud.Project

	oc := metrics.ObserveCall("GetProject")
	err := p.client.GetWithContext(ctx, "/cloud/project/"+projectId, &project)
	oc.End(err)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (p *node) Loadbalancer(project string) loadbalancer.Node {
	return loadbalancer.New(p.client, project)
}

func (p *node) Capabilities(project string) capabilities.Capabilities {
	return capabilities.New(p.client, project)
}
