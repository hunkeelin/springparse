package springparse

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	putSuccess = promauto.NewCounter(prometheus.CounterOpts{
		Name: "springparse_putsuccess",
		Help: "The total number of index created",
	})
	putFailed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "springparse_putfailed",
		Help: "The total number of index failed to put",
	})
)
