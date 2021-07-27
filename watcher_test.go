package main

import (
	"circuitbreaker/config"
	"circuitbreaker/metric_collector"
	"github.com/stretchr/testify/assert"
	"testing"
)

type fakeHttpClient struct {
	fakeMessage string
}

func (fh *fakeHttpClient) Get(url string) (metric_collector.HttpResponse, error) {
	return metric_collector.HttpResponse{
		StatusCode: 200,
		Body: fh.fakeMessage,
	}, nil
}

var normalTempMetric =
	"node_fake_some_elem{device=\"sda3\"}\nnode_hwmon_temp_celsius{chip=\"platform_coretemp_1\",sensor=\"temp2\"} 90\n"

func TestProcessMetric_WHEN_허용온도를_넘은_장비가_없는_경우에_THEN_서킷브레이커가_작동하지_않는다(t *testing.T) {
	lowTriggerPointWatcher := NewWatcher(&config.CpuConfig{
		TriggerPoint:    100,
		TolerableNumber: 0,
	}, &config.GpuConfig{
		TriggerPoint:    100,
		TolerableNumber: 0,
	}, []string{"192.168.1.11:9100"}, []string{}, nil,
		metric_collector.NewHttpCollector(&fakeHttpClient{fakeMessage: normalTempMetric}))

	braked, repaired := lowTriggerPointWatcher.ProcessMetric()
	assert.Equal(t, 0, len(braked))
	assert.Equal(t, 0, len(repaired))
}

func TestProcessMetric_WHEN_overheated_devices_exists_RETURNS_braked_list(t *testing.T) {
	lowTriggerPointWatcher := NewWatcher(&config.CpuConfig{
		TriggerPoint:    10,
		TolerableNumber: 0,
	}, &config.GpuConfig{
		TriggerPoint:    10,
		TolerableNumber: 0,
	}, []string{"192.168.1.11:9100"}, []string{}, nil,
		metric_collector.NewHttpCollector(&fakeHttpClient{fakeMessage: normalTempMetric}))

	braked, repaired := lowTriggerPointWatcher.ProcessMetric()
	assert.Equal(t, 1, len(braked))
	assert.Equal(t, 0, len(repaired))
}

/*
func TestShouldCircuitBreakingNodeExporterMetric_WHEN_overheated_cpu_is_more_than_tolerable_num_THAN_True(t *testing.T) {
	watcher := NewWatcher(&CcbConfig{
		CpuTriggerPoint: 1,
		CpuTolerableNumber: 1,
		NodeExporters: []string{"127.0.0.1:9100"},
	})
	result := watcher.shouldCircuitBreakingNodeExporterMetric("localhost:9100", fakeNodeExporterMetric)
	assert.Equal(t, len(result), 1)
}
func TestShouldCircuitBreakingNodeExporterMetric_WHEN_cpu_temp_lower_than_trigger_point_THAN_False(t *testing.T) {
	watcher := NewWatcher(&CcbConfig{
		CpuTriggerPoint: math.MaxInt32,
		CpuTolerableNumber: math.MaxInt32,
		NodeExporters: []string{"127.0.0.1:9100"},
	})
	result := watcher.shouldCircuitBreakingNodeExporterMetric("localhost:9100", fakeNodeExporterMetric)
	assert.Equal(t, len(result), 0)
}
func TestShouldCircuitBreakingNodeExporterMetric_WHEN_overheated_cpu_smaller_than_tolerable_THAN_False(t *testing.T) {
	watcher := NewWatcher(&CcbConfig{
		CpuTriggerPoint: 0,
		CpuTolerableNumber: math.MaxInt32,
		NodeExporters: []string{"127.0.0.1:9100"},
	})
	result := watcher.shouldCircuitBreakingNodeExporterMetric("localhost:9100", fakeNodeExporterMetric)
	assert.Equal(t, len(result), 0)
}
*/