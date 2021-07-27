package metric

import (
	"circuitbreaker/config"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

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
	ShouldBreak bool
	ShouldRepair bool
}

func NewMetric(cpuConfig *config.CpuConfig, gpuConfig *config.GpuConfig, rawMetric *ExporterMetric) (Metric, error) {
	metricContent := rawMetric.Metric

	lines := strings.Split(metricContent, "\n")

	for _, line := range lines {
		if strings.Contains(line, "node_hwmon_temp_celsius") {
			toks := strings.Split(line, " ")
			temp, err :=  strconv.Atoi(toks[1])
			if err != nil {
				log.Errorf("올바르지 않은 node exporter 메트릭이 수집되었습니다! err = [%s], metric = [%s]",
					err.Error(), metricContent)
				return &NodeExporterMetrics{}, err
			}
			return &NodeExporterMetrics{
				temp,
				cpuConfig,
			}, nil
		}
	}


	if strings.Contains(metricContent, "DCGM_FI_DEV_GPU_TEMP") {
		return &DcgmExporterMetrics{}, nil
	}
	return &DcgmExporterMetrics{}, nil
}