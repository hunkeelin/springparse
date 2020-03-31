package springparse

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	putSuccess = promauto.NewCounter(prometheus.CounterOpts{
		Name: "putsuccess",
		Help: "The total number of index created",
	})
)
