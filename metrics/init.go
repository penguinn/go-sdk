package metrics

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 在应用初始化的时候调用此接口
func Init(g *gin.Engine, namespace, systemName string) error {
	err := initOpts(namespace, systemName)
	if err != nil {
		return err
	}

	g.Use(MetricMiddleware)

	g.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return nil
}

func initOpts(namespace, systemName string) error {
	defaultMetricManager = &MetricManager{
		systemName: systemName,
		metrics:    sync.Map{},
	}
	if namespace != "" {
		defaultMetricManager.namespace = namespace
	} else {
		defaultMetricManager.namespace = DefaultNamespace
	}
	err := initDefaultMetrics()
	if err != nil {
		return err
	}

	return nil
}

func initDefaultMetrics() error {
	labels := []string{Method, Path, Status}

	_, err := GenerateCounter(&MetricValue{
		ValueType: Counter,
		Name:      ReqCount,
		Help:      "Counter. Total number of HTTP requests made",
		Labels:    labels,
	})
	if err != nil {
		return err
	}

	_, err = GenerateHistogram(&MetricValue{
		ValueType: Histogram,
		Name:      ReqDuration,
		Help:      "Histogram. HTTP request latencies in seconds",
		Labels:    labels,
	})
	if err != nil {
		return err
	}

	_, err = GenerateSummary(&MetricValue{
		ValueType: Summary,
		Name:      ReqSizeBytes,
		Help:      "Summary. HTTP request sizes in bytes",
		Labels:    labels,
	})
	if err != nil {
		return err
	}

	_, err = GenerateSummary(&MetricValue{
		ValueType: Summary,
		Name:      RespSizeBytes,
		Help:      "Summary. HTTP request sizes in bytes",
		Labels:    labels,
	})
	if err != nil {
		return err
	}

	return nil
}

// 初始化metric指标
func InitMetric(metricsValues []*MetricValue) error {
	for _, metricsValue := range metricsValues {
		if f, ok := promTypeHandler[metricsValue.ValueType]; ok {
			_, err := f(metricsValue)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("bml-go-sdk metrics init error, unknown valuetype:%s", metricsValue.ValueType)
		}
	}

	return nil
}
