package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func MetricMiddleware(c *gin.Context) {
	start := time.Now()
	c.Next()

	method := c.Request.Method
	path := c.Request.URL.Path
	status := strconv.Itoa(c.Writer.Status())

	labelValues := map[string]string{Method: method, Path: path, Status: status}

	metric, err := defaultMetricManager.GetMetric(ReqCount)
	if err == nil && metric != nil {
		_ = metric.IncWithLabel(labelValues)
	}

	metric, err = defaultMetricManager.GetMetric(ReqSizeBytes)
	if err == nil && metric != nil {
		_ = metric.ObserveWithLabel(labelValues, calcRequestSize(c.Request))
	}

	metric, err = defaultMetricManager.GetMetric(ReqSizeBytes)
	if err == nil && metric != nil {
		_ = metric.ObserveWithLabel(labelValues, float64(c.Writer.Size()))
	}

	metric, err = defaultMetricManager.GetMetric(ReqDuration)
	if err == nil && metric != nil {
		_ = metric.ObserveWithLabel(labelValues, time.Since(start).Seconds())
	}
}

func calcRequestSize(r *http.Request) float64 {
	size := 0
	if r.URL != nil {
		size = len(r.URL.String())
	}

	size += len(r.Method)
	size += len(r.Proto)

	for name, values := range r.Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}
	size += len(r.Host)

	// r.Form and r.MultipartForm are assumed to be included in r.URL.
	if r.ContentLength != -1 {
		size += int(r.ContentLength)
	}
	return float64(size)
}
