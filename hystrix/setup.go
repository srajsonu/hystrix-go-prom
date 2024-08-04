package hystrix

import (
	hystrixmetric "github.com/afex/hystrix-go/hystrix/metric_collector"
	"github.com/prometheus/client_golang/prometheus"
)

func SetupHystrix(namespace string, configs map[string]Config) {
	initHystrixCommandConfigs(configs)
	registerHystrixPromMetricsCollector(namespace)
}

func initHystrixCommandConfigs(configs map[string]Config) {
	for command, config := range configs {
		setupHystrixConfig(command, config)
	}
}

func registerHystrixPromMetricsCollector(namespace string) {
	wrapper := NewPrometheusCollector(namespace, nil, prometheus.DefBuckets)
	// register and initialize to hystrix prometheus
	hystrixmetric.Registry.Register(wrapper.Collector)
}
