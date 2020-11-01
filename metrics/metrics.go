package metrics

// metrics package expose a set of metrics filled by the SDK about API calls

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace = "ovh"
	Subsystem = "sdk"
)

var (
	Registry = prometheus.NewRegistry()

	CallCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: Subsystem,
		Name:      "call_count",
		Help:      "OVH API call count",
	}, []string{"action"})

	ErroredCallCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: Subsystem,
		Name:      "call_error_count",
		Help:      "OVH API errored call count",
	}, []string{"action"})

	CallDuration = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: Subsystem,
		Name:      "call_duration",
		Help:      "OVH API call duration",
		ConstLabels: prometheus.Labels{
			"unit": "ms",
		},
	}, []string{"action"})
)

func init() {
	Registry.MustRegister(
		CallCount,
		CallDuration,
	)
}
