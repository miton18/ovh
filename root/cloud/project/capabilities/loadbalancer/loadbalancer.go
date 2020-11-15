package loadbalancer

import (
	ovhclient "github.com/ovh/go-ovh/ovh"

	"github.com/miton18/ovh/root/cloud/project/capabilities/loadbalancer/region"
)

type Loadbalancer interface {
	Region() region.Region
}

type loadbalancer struct {
	client *ovhclient.Client
	project string
}

func New(client *ovhclient.Client, project string) Loadbalancer {
	return &loadbalancer{client, project}
}

func (l *loadbalancer) Region() region.Region {
	return region.New(l.client, l.project)
}
