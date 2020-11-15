# OVH SDK

[![PkgGoDev](https://pkg.go.dev/badge/github.com/clever-telemetry/ovh)](https://pkg.go.dev/github.com/clever-telemetry/ovh)
![License](https://img.shields.io/github/license/clever-telemetry/ovh?style=flat)

Full featured OVH SDK

This project heavily rely on auto-generated [OVH API models](https://github.com/clever-telemetry/ovh-models) source code repository.

## Getting started

### Install the SDK

```sh
go get -u github.com/clever-telemetry/ovh
```

### Keys

To contact OVH API, we need 3 keys, `applicationKey`, `applicationSecret`, `consumerKey`
You can craft them on the official page: [CreateToken](https://api.ovh.com/createToken/)

### First call

Make our first call, let's try to list all our cloud projects

```golang
package main

import (
	"context"
	"fmt"

	"github.com/clever-telemetry/ovh"
)

func main() {
	client, err := ovh.New(ovh.OVH_EU, "<APPLICATION_KEY>", "<APPLICATION_SECRET>", "<CONSUMER_KEY>")
	if err != nil {
		panic(err.Error())
	}

    cloudProjectList, err := client.Cloud().Project().List(context.Background())
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Projects %#v", cloudProjectList)
}

```

## Contribute

### The method you want is not implemented
You can request a new method by creating a new [method issue](https://github.com/clever-telemetry/ovh/issues/new?assignees=miton18&labels=enhancement%2C+method+request&template=method-request.md&title=)

### You want to code a new method ? 
See [CONTRIBUTING](./CONTRIBUTING.md)

## Tips

## Observability
Each SDK action is observed, so you can find a list of metrics in the [metrics](https://github.com/clever-telemetry/ovh/tree/master/metrics) package.
This is up to you to register them.

TODO: Tracing

### Examples

#### Get a loadbalancer

```golang
package main

import (
	"context"
	"fmt"

	"github.com/clever-telemetry/ovh"
)

func main() {
	client := ovh.MustNewFromEnv()

	loadbalancer, err := client.
	    Cloud().
	    Project().
	    Loadbalancer("<PROJECT_ID>").
	    Get(context.Background(), "<LOADBALANCER_ID>")
	

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("RES %#v", test)
}
```
