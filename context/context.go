package context

import (
	"context"

	ovhclient "github.com/ovh/go-ovh/ovh"
)

type clientKey struct{}

func WithClient(ctx context.Context, client *ovhclient.Client) context.Context {
	return context.WithValue(ctx, clientKey{}, client)
}

func Client(ctx context.Context) *ovhclient.Client {
	return ctx.Value(clientKey{}).(*ovhclient.Client)
}

type projectIdKey struct{}

func WithProjectId(ctx context.Context, projectId string) context.Context {
	return context.WithValue(ctx, projectIdKey{}, projectId)
}

func ProjectId(ctx context.Context) string {
	return ctx.Value(projectIdKey{}).(string)
}

type loadbalancerIdKey struct{}

func WithLoadbalancerId(ctx context.Context, loadbalancerId string) context.Context {
	return context.WithValue(ctx, loadbalancerIdKey{}, loadbalancerId)
}

func LoadbalancerId(ctx context.Context) string {
	return ctx.Value(loadbalancerIdKey{}).(string)
}
