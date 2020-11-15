package config

type Endpoint string

var (
	OVH_EU Endpoint = "ovh-eu"
	OVH_CA Endpoint = "ovh-ca"
	OVH_US Endpoint = "ovh-us"
)

func (ep Endpoint) String() string {
	return string(ep)
}
