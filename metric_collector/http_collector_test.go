package metric_collector

import (
	"circuitbreaker/metric"
	"github.com/stretchr/testify/assert"
	"testing"
)

type fakeHttpClient struct {
	fakeMessage string
}

func (fh *fakeHttpClient) Get(url string) (httpResponse, error) {
	return httpResponse{
		statusCode: 200,
		body: fh.fakeMessage,
	}, nil
}

func TestCollect(t *testing.T) {
	endpoint := "192.168.1.11:9100"
	fakeRespBody := `node_hwmon_temp_celsius{chip=\"platform_coretemp_1\",sensor=\"temp2\"} 90\n`
	coll := NewHttpCollector(&fakeHttpClient{fakeRespBody})
	resultchan := make(chan *metric.ExporterMetric)
	go coll.Collect(resultchan, endpoint)

	//results := []*metric.ExporterMetric{}
	go func() {
		for {
			select {
			case collectResult := <- resultchan:
				assert.Equal(t, endpoint, collectResult.MetricHost)
				assert.Equal(t, fakeRespBody, collectResult.Metric)
			}
		}
	}()
}
