package ovh

import (
	ovhclient "github.com/ovh/go-ovh/ovh"
)

type Endpoint string

var (
	OVH_EU        Endpoint = ovhclient.OvhEU
	OVH_CA        Endpoint = ovhclient.OvhCA
	OVH_US        Endpoint = ovhclient.OvhUS
	KIMSUFI_EU    Endpoint = ovhclient.KimsufiEU
	KIMSUFI_CA    Endpoint = ovhclient.KimsufiCA
	SOYOUSTART_EU Endpoint = ovhclient.SoyoustartEU
	SOYOUSTART_CA Endpoint = ovhclient.SoyoustartCA
)

func (ep Endpoint) String() string {
	return string(ep)
}
