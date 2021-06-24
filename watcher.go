package main

import (
	"circuitbreaker/circuitbreaker"
	"circuitbreaker/metric"
	"circuitbreaker/metric_collector"
	"database/sql"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"sync"
)

type Watcher struct {
	mutex *sync.Mutex
	cpuConfig *CpuConfig
	gpuConfig *GpuConfig
	nodeExporters []string
	dcgmExporters []string
	collector metric_collector.MetricCollector
	circuitBreaker circuitbreaker.CircuitBreaker
	db *sql.DB
	cpuTriggerPoint int
	gpuTriggerPoint int
	brakedHostCache []string
	restClient *resty.Client
}

const (
	CPU_BRAKE = "CPU_BRAKE"
	GPU_BRAKE = "GPU_BRAKE"
)

func NewWatcher(cpuConfig *CpuConfig, gpuConfig *GpuConfig, nodeExporters []string, dcgmExporters []string, db *sql.DB, mc metric_collector.MetricCollector) *Watcher {

	return &Watcher{
		mutex: &sync.Mutex{},
		cpuConfig: cpuConfig,
		gpuConfig: gpuConfig,
		nodeExporters: nodeExporters,
		dcgmExporters: dcgmExporters,
		collector: mc,
		db: db,
		restClient: resty.New(),
	}
}

func (w *Watcher) run() {
	for ;; {

	}
}

var NODE_HOWMON_TEMP_CELSIUS string = "node_hwmon_temp_celsius"

func (w *Watcher) sendCircuitBreakingAlert() {

}

func (w *Watcher) shouldCircuitBreakingNodeExporterMetric(hostname, metricStr string) ([]string) {
	metrics := strings.Split(metricStr, "\n")
	tolerableCount := 0
	var overheated []string
	for _, metric := range metrics {
		//fmt.Println(metric)
		if strings.Contains(metric, NODE_HOWMON_TEMP_CELSIUS) {
			splitted := strings.Split(metric, " ")
			if len(splitted) < 2 {
				log.Errorf("%s 의 온도 정보가 올바르지 않습니다.", hostname)
				continue
			}
			temp, err :=  strconv.Atoi(splitted[1])
			if err != nil {
				log.Errorf("%s 의 온도 정보의 포맷이 올바르지 않습니다.", hostname)
				continue
			}
			if  temp >= w.cpuTriggerPoint {
				tolerableCount++
				overheated = append(overheated, metric)
			}
		}
	}
	if tolerableCount >= w.cpuConfig.tolerableNumber {
		return overheated
	}

	if tolerableCount == 0 {
		_, err := w.db.Exec("DELETE FROM BRAKED_HOST WHERE hostname=? AND brake_type=?", hostname, CPU_BRAKE)
		if err != nil {
			return []string{}
		}
		_, err = w.db.Exec("INSERT INTO CB_LOG (hostname, brake_type, braked) VALUES (?, ?)", hostname, CPU_BRAKE, 0)
		if err != nil {
			log.Errorf("Failed to write CB_LOG : %s", err.Error())
		}
	}

	return []string{}
}

func (w *Watcher) fetchNodeExporterMetric(channel <- chan ExporterMetric, metricHost string) {


}

func (w *Watcher) collectMetric(metricChan <- chan *metric.ExporterMetric, metricEndpoints []string) {
	for _, metricHost := range metricEndpoints {
		go w.collector.Collect(metricChan, metricHost)
	}
}

func (w *Watcher) ProcessMetric() {
	rawMetricChan := make(chan *metric.ExporterMetric)
	merged := append(w.nodeExporters, w.dcgmExporters...)

	w.collectMetric(rawMetricChan, merged)

	for {
		select {
			case rawMetric := <- rawMetricChan:
				concreteMetric := metric.NewMetric(rawMetric)
				rst := concreteMetric.Diagnose()
				w.circuitBreaker.Break(rst.MetricType, rst.CircuitBreakCandidates)
				w.circuitBreaker.Repair(rst.MetricType, rst.CircuitRecoveryCandidates)
		}
	}
}
