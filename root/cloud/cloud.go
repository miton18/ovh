package cloud

import (
	ovhclient "github.com/ovh/go-ovh/ovh"

	"github.com/miton18/ovh/root/cloud/project"
)

type Node interface {
	Project() project.Node
}

type node struct {
	client *ovhclient.Client
}

func New(client *ovhclient.Client) Node {
	return &node{client}
}

func (c *node) Project() project.Node {
	return project.New(c.client)
}
