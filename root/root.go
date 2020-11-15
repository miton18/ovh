package root

import (
	"github.com/miton18/ovh/root/me"
	ovhclient "github.com/ovh/go-ovh/ovh"

	"github.com/miton18/ovh/root/cloud"
)

type Node interface {
	Cloud() cloud.Node
	Me() me.Node
}

type node struct {
	client *ovhclient.Client
}

func New(client *ovhclient.Client) Node {
	return &node{client}
}

func (c *node) Cloud() cloud.Node {
	return cloud.New(c.client)
}

func (c *node) Me() me.Node {
	return me.New(c.client)
}
