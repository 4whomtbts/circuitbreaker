package metric_collector

import (
	"circuitbreaker/metric"
	log "github.com/sirupsen/logrus"
)

type httpResponse struct {
	statusCode int
	body string
}

type httpClient interface {
	Get(url string) (httpResponse, error)
}

type HttpCollector struct {
	client httpClient
}

func NewHttpCollector (client httpClient) *HttpCollector {
	return &HttpCollector{
		client: client,
	}
}

func (hc *HttpCollector) Collect(metricChan chan *metric.ExporterMetric, metricEndpoint string) {
	resp, err  := hc.client.Get(metricEndpoint)
	if err != nil {
		log.Errorf("메트릭 엔드포인트 [ %s ] 로 부터 응답을 얻는데 실패했습니다 : %s", metricEndpoint, err.Error())
		return
	}
	metricChan <- metric.NewExporterMetric(metricEndpoint, resp.body)
}
