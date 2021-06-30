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

func (sb *SoftCircuitBreaker) Break(metricType string, host string) bool {
	fmt.Println("softCircuitBreaker called")
	return true
}

func (sb *SoftCircuitBreaker) Repair(metricType string, host string) bool {
	fmt.Println("softCircuitBreaker repaired!")
	return true
}

