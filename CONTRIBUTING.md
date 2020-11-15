# Contributing

You want to add a missing feature to the OVH SDK ?
Depending of the feature you can either implement a [new method on an existing node](#write-a-new-method),
or [add a new node](#write-a-new-node).


## Write a new method

First, lookup the right node for your method.
For example, if I want to implement a method to get a cloud project, I will look at `./root/cloud/project/project.go`

You can see this file own a node, which is an interface, as a SDK user, you only interfaced client, so let's add a new method on it.

```golang
type Client interface {
	List(ctx context.Context) ([]string, error)
+	Get(ctx context.Context, projectId string) (*cloud.Project, error)
	Loadbalancer(string) loadbalancer.Client
    Capabilities(string) capabilities.Client
}
```

Rules:
- All direct methods must start with a `Context`, next to the arguments you need, here, the `projectId`
- All direct methods must returns an error, and optionally a response. when response is a struct, use a pointer to it, here, `*cloud.Project`

You can now, implement the call you define in the interface:

```golang
func (c *client) Get(ctx context.Context, projectId string) (*cloud.Project, error) {
    return nil, nil
}
```

Use the embeded OVH API client to perform calls:

```golang
func (c *client) Get(ctx context.Context, projectId string) (*cloud.Project, error) {
	var project cloud.Project

	err := c.client.GetWithContext(ctx, "/cloud/project/"+projectId, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
```

Finally, add observability to your call, all SDK methods need to be observable, you can use a helper to do that:

```golang
func (c *client) Get(ctx context.Context, projectId string) (*cloud.Project, error) {
	var project cloud.Project

+	oc := metrics.ObserveCall("GetProject")
	err := c.client.GetWithContext(ctx, "/cloud/project/"+projectId, &project)
+	oc.End(err)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
```

The observability context take an action as parameter, your action must be UpperCamelCased, and unique to your method.
When you call `ObservabilityContext.End(error)`, we compute the duration between the creation of the context and the `End()` call,
so you must be as near as possible of the IO time consuming method.

Your new Method is ready !

## Write a new node

Each OVH API pars is translated to a node, for example `/cloud/project` owns 2 nodes, one for `Cloud` and another for `Project`.
You can add a node at the level you want, but it must translate a real API tree structure.

Let's add a new node for Loadbalancers under Cloud projects:

```golang
// /root/cloud/project/loadbalancer/loadbalancer.go
package loadbalancer

import (
    "context"
	"fmt"
	"github.com/clever-telemetry/ovh-models/cloud"
	self "github.com/clever-telemetry/ovh/context"
	"github.com/clever-telemetry/ovh/metrics"
	ovhclient "github.com/ovh/go-ovh/ovh"
	"github.com/clever-telemetry/ovh/utils"
)

type Client interface {
	List(ctx context.Context) ([]string, error)
	Get(ctx context.Context, loadbalancerId string) (*cloud.LoadBalancer, error)
	Create(ctx context.Context, region, name, description string) (*cloud.LoadBalancer, error)
	SetName(ctx context.Context, loadbalancerId, name string) (*cloud.LoadBalancer, error)
	SetDescription(ctx context.Context, loadbalancerId, description string) (*cloud.LoadBalancer, error)
	Delete(ctx context.Context, loadbalancerId string) error
}

type client struct {
	ctx       context.Context
	client    *ovhclient.Client
	projectId string
}

```

Here, we define a new `Client{}` interface for out node, and a corresponding implementation with `client{}`.
Our `client{}` will owns all needed informations perform `Client` interface methods.
So, basicly, we need an OVH API client to perform requests, with `self.Client(ctx)` we can get things from context.
If you need to add required informations, you can add them to the local `context` package.

Cloud Project ID is also mandatory, so we add `self.ProjectId(ctx)`.

Rule: Always keep the current context in your node.

Then, we add a `New(Context)` method:

```golang
package loadbalancer

...

func New(ctx context.Context) Client {
	return &client{
		ctx:       ctx,
		client:    self.Client(ctx),
		projectId: self.ProjectId(ctx),
	}
}
```
 
This standard method only take a `Context` as parameter and return a `Client{}` interface.

We can now add our node to the parent one:

```golang
// /root/cloud/project/project.go
type Client interface {
	List(context.Context) ([]string, error)
	Get(context.Context, string) (*cloud.Project, error)
+	Loadbalancer(string) loadbalancer.Client
	Capabilities(string) capabilities.Client
}
```

And implement this method:

```golang
func (c *client) Loadbalancer(project string) loadbalancer.Client {
	ctx := self.WithProjectId(c.ctx, project) // we add loadbalancer required informations
	return loadbalancer.New(ctx)
}
```

Your new node is ready !