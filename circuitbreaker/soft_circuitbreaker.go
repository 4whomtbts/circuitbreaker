package circuitbreaker

import (
	"fmt"
	"sync"
)

type SoftCircuitBreaker struct {
	mutex *sync.Mutex
	sshUser string
	sshPwd string
}

func NewSoftCircuitBreaker(sshUser, sshPwd string) *SoftCircuitBreaker {
	return &SoftCircuitBreaker{
		sshUser: sshUser,
		sshPwd: sshPwd,
	}
}

func (sb *SoftCircuitBreaker) Break(metricType string, hosts []string) {
	fmt.Println("softCircuitBreaker called")
}

func (sb *SoftCircuitBreaker) Repair(metricType string, hosts []string) {
	fmt.Println("softCircuitBreaker repaired!")
}

