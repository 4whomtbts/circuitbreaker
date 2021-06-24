package circuitbreaker

type CircuitBreaker interface {
	Break(metricType string, hosts []string)
	Repair(metricType string, hosts []string)
}
