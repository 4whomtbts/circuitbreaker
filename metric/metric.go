package metric

import "strings"

type Metric interface {
	Diagnose() MetricDiagnosisResult
}

type ExporterMetric struct {
	MetricHost string
	Metric string
}

func NewExporterMetric(metricHost, metric string) *ExporterMetric {
	return &ExporterMetric{
		MetricHost: metricHost,
		Metric:     metric,
	}
}

type MetricDiagnosisResult struct {
	MetricType string
	CircuitBreakCandidates []string
	CircuitRepairCandidates []string
}

func NewMetric(rawMetric *ExporterMetric) Metric {
	metricCont := rawMetric.Metric
	if strings.Contains(metricCont, "node_hwmon_temp_celsius") {
		return &NodeExporterMetrics{}
	}
	return &NodeExporterMetrics{}
}
