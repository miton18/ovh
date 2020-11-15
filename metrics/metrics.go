package metrics

// metrics package expose a set of metrics filled by the SDK about API calls

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type ObservabilityContext struct {
	start  time.Time
	action string
}

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

// observe an HTTP call, action must be UpperCamelCased
func ObserveCall(action string) *ObservabilityContext {
	return &ObservabilityContext{
		start:  time.Now(),
		action: action,
	}
}

// End the observation with a nillable error
func (ctx *ObservabilityContext) End(err error) {
	d := time.Since(ctx.start)

	if err != nil {
		CallCount.WithLabelValues(ctx.action).Inc()
	} else {
		ErroredCallCount.WithLabelValues(ctx.action).Inc()
	}

	CallDuration.
		WithLabelValues(ctx.action).
		Add(float64(d.Milliseconds()))
}
