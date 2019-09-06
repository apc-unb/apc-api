package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	InvocationCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "invocation_count",
		Help: "The total number of requests receivid",
	})

	Uptime = promauto.NewCounter(prometheus.CounterOpts{
		Name: "uptime",
		Help: "The total time that the server is up",
	})
)
