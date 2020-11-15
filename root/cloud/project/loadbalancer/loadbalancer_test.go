package loadbalancer

import (
	"context"
	"os"
	"testing"

	"github.com/clever-telemetry/ovh-models/cloud"
	self "github.com/clever-telemetry/ovh/context"
	"github.com/clever-telemetry/ovh/root/cloud/project/loadbalancer/configuration"
	ovhclient "github.com/ovh/go-ovh/ovh"
)

func TestIntegration(t *testing.T) {
	_client, err := ovhclient.NewClient(
		os.Getenv("OVH_ENDPOINT"),
		os.Getenv("OVH_APP_KEY"),
		os.Getenv("OVH_APP_SECRET"),
		os.Getenv("OVH_CONSUMER_KEY"),
	)

	cctx := context.Background()
	cctx = self.WithClient(cctx, _client)
	cctx = self.WithProjectId(cctx, os.Getenv("OVH_PROJECT"))

	client := New(cctx)
	ctx := context.Background()

	list, err := client.List(ctx)
	if err != nil {
		t.Errorf("can't list Loadbalancers: %s", err.Error())
	}
	t.Logf("List %+v", list)

	lb, err := client.Create(ctx, "GRA", "sdk-unit-test", "for unit tests purpose (deletable)")
	if err != nil {
		t.Errorf("can't create Loadbalancers: %s", err.Error())
	}
	t.Logf("Loadbalancer ID: %s", lb.Id)

	// Delete lb after tests
	t.Cleanup(func() {
		client.Delete(context.Background(), lb.Id)
	})

	_, err = client.SetName(ctx, lb.Id, "sdk-unit-test-updated")
	if err != nil {
		t.Errorf("can't update loadbalancer: %s", err.Error())
	}

	_, err = client.SetDescription(ctx, lb.Id, "for unit tests purpose (deletable)")
	if err != nil {
		t.Errorf("can't update loadbalancer: %s", err.Error())
	}

	// CONFIGURATION
	cctx = self.WithLoadbalancerId(cctx, lb.Id)
	configClient := configuration.New(cctx)

	config, err := configClient.Create(ctx, &cloud.ConfigurationCreation{})
	if err != nil {
		t.Errorf("can't create loadbalancer configuration: %s", err.Error())
	}

	_, err = configClient.List(ctx)
	if err != nil {
		t.Errorf("can't list loadbalancer configuration: %s", err.Error())
	}

	_, err = configClient.Get(ctx, config.Version)
	if err != nil {
		t.Errorf("can't get loadbalancer configuration: %s", err.Error())
	}

	err = configClient.Apply(ctx, config.Version)
	if err != nil {
		t.Errorf("can't apply loadbalancer configuration: %s", err.Error())
	}

	// we need a non applied config to delete
	config2, err := configClient.Create(ctx, &cloud.ConfigurationCreation{})
	if err != nil {
		t.Errorf("can't create loadbalancer configuration2: %s", err.Error())
	}
	err = configClient.Delete(ctx, config2.Version)
	if err != nil {
		t.Errorf("can't delete loadbalancer configuration2: %s", err.Error())
	}

	err = client.Delete(ctx, lb.Id)
	if err != nil {
		t.Errorf("can't delete Loadbalancers: %s", err.Error())
	}
}
