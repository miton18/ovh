package me

import (
	"context"
	models "github.com/miton18/ovh-models/me"
	"github.com/miton18/ovh/metrics"
	ovhclient "github.com/ovh/go-ovh/ovh"
)

type Node interface {
	Get(ctx context.Context) (*models.Nichandle, error)
}

type node struct {
	client *ovhclient.Client
}

func New(client *ovhclient.Client) Node {
	return &node{client}
}

func (c *node) Get(ctx context.Context) (*models.Nichandle, error) {
	var o models.Nichandle

	oc := metrics.ObserveCall("ListProjects")
	err := c.client.GetWithContext(ctx, "/me", &o)
	oc.End(err)

	return &o, err
}
