package circuitbreaker

type CircuitBreaker interface {
	Break(metricType string, host string) bool
	Repair(metricType string, host string) bool
}
