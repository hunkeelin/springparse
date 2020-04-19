package springparse

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	putSuccess = promauto.NewCounter(prometheus.CounterOpts{
		Name: "springparse_putsuccess",
		Help: "The total number batch send due to buffer overload",
	})
	putFailed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "springparse_putfailed",
		Help: "The total number of index failed to put",
	})
	putFlushSuccess = promauto.NewCounter(prometheus.CounterOpts{
		Name: "springparse_putflushsuccess",
		Help: "The total of batch sent due to time flushes",
	})
)
