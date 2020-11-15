package capabilities

import (
	ovhclient "github.com/ovh/go-ovh/ovh"

	"github.com/miton18/ovh/root/cloud/project/capabilities/loadbalancer"
)

type Capabilities interface {
	Loadbalancer() loadbalancer.Loadbalancer
}

type capabilities struct {
	client *ovhclient.Client
	project string
}

func New(client *ovhclient.Client, project string) Capabilities {
	return &capabilities{client, project}
}

func (c *capabilities) Loadbalancer() loadbalancer.Loadbalancer {
	return loadbalancer.New(c.client, c.project)
}
