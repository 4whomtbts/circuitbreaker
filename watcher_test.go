package main

var fakeNodeExporterMetric =
"node_fake_some_elem{device=\"sda3\"}\nnode_hwmon_temp_celsius{chip=\"platform_coretemp_1\",sensor=\"temp2\"} 90\n"

/*
func TestProcessMetric(t *testing.T) {
	w := NewWatcher(&CpuConfig{
		triggerPoint:    10,
		tolerableNumber: 0,
	}, &GpuConfig{
		triggerPoint:    10,
		tolerableNumber: 0,
	}, []string{}, []string{}, nil, &metric_collector.HttpCollector{})

	w.ProcessMetric()
}
*/
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