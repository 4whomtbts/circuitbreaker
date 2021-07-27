package metric

import "circuitbreaker/config"

type NodeExporterMetrics struct {
	temp int
	cpuConfig *config.CpuConfig
}

func (nem *NodeExporterMetrics) GetCircuitBreakRequiredHosts() []string {
	return []string{}
}

func (nem *NodeExporterMetrics) GetCircuitRecoveryRequiredHosts() []string {
	return []string{}
}

func (nem *NodeExporterMetrics) Diagnose() MetricDiagnosisResult {

	if nem.cpuConfig.TriggerPoint <= nem.temp {
		return MetricDiagnosisResult{
			MetricType: "NODE_EXPORTER",
			ShouldBreak: true,
			ShouldRepair: false,
		}
	}
	return MetricDiagnosisResult{}
}

/*
diagnose
 - BreakReq
 - RecReq
break(BreakReq)
recover(REcReq)
 */