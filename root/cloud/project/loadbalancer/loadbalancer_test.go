package loadbalancer

import (
	"context"
	"os"
	"testing"

	"github.com/miton18/ovh-models/cloud"
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
	t.Logf("Node ID: %s", lb.Id)

	// Delete lb after tests
	t.Cleanup(func() {
		client.Delete(context.Background(), lb.Id)
	})

	_, err = client.SetName(ctx, lb.Id, "sdk-unit-test-updated")
	if err != nil {
		t.Errorf("can't update node: %s", err.Error())
	}

	_, err = client.SetDescription(ctx, lb.Id, "for unit tests purpose (deletable)")
	if err != nil {
		t.Errorf("can't update node: %s", err.Error())
	}

	// CONFIGURATION
	config, err := configClient.Create(ctx, &cloud.ConfigurationCreation{})
	if err != nil {
		t.Errorf("can't create node configuration: %s", err.Error())
	}

	_, err = configClient.List(ctx)
	if err != nil {
		t.Errorf("can't list node configuration: %s", err.Error())
	}

	_, err = configClient.Get(ctx, config.Version)
	if err != nil {
		t.Errorf("can't get node configuration: %s", err.Error())
	}

	err = configClient.Apply(ctx, config.Version)
	if err != nil {
		t.Errorf("can't apply node configuration: %s", err.Error())
	}

	// we need a non applied config to delete
	config2, err := configClient.Create(ctx, &cloud.ConfigurationCreation{})
	if err != nil {
		t.Errorf("can't create node configuration2: %s", err.Error())
	}
	err = configClient.Delete(ctx, config2.Version)
	if err != nil {
		t.Errorf("can't delete node configuration2: %s", err.Error())
	}

	err = client.Delete(ctx, lb.Id)
	if err != nil {
		t.Errorf("can't delete Loadbalancers: %s", err.Error())
	}
}
