package metric_collector

import "circuitbreaker/metric"

type MetricCollector interface {
	Collect(metricChan chan *metric.ExporterMetric, metricEndpoint string)
}
