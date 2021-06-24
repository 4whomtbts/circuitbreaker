package metric

type NodeExporterMetrics struct {

}

func (nem *NodeExporterMetrics) GetCircuitBreakRequiredHosts() []string {
	return []string{}
}

func (nem *NodeExporterMetrics) GetCircuitRecoveryRequiredHosts() []string {
	return []string{}
}

func (nem *NodeExporterMetrics) Diagnose() MetricDiagnosisResult {

	return MetricDiagnosisResult{

	}
}

/*
diagnose
 - BreakReq
 - RecReq
break(BreakReq)
recover(REcReq)
 */