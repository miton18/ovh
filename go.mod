module github.com/miton18/ovh

go 1.15

replace github.com/miton18/ovh-models => ../ovh-models

require (
	github.com/miton18/ovh-models v0.0.0-20210624141038-ace9ef160787
	github.com/ovh/go-ovh v1.1.0
	github.com/prometheus/client_golang v1.8.0
	github.com/prometheus/common v0.15.0 // indirect
	github.com/spf13/viper v1.8.0
)
