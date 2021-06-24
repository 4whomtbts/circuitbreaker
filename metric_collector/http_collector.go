package metric_collector

import "circuitbreaker/metric"

type HttpCollector struct {

}

func (hc *HttpCollector) Collect(metricChan <- chan *metric.ExporterMetric, metricEndpoint string) {

}
