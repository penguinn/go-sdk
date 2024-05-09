package metrics

const (
	DefaultNamespace = "bml"
)

type MetricType string

const (
	Counter   MetricType = "counter"
	Gauge     MetricType = "gauge"
	Histogram MetricType = "histogram"
	Summary   MetricType = "summary"
	Untyped   MetricType = "untyped"
)

const (
	ReqCount      = "http_request_count"
	ReqDuration   = "http_request_duration_seconds"
	ReqSizeBytes  = "http_request_size_bytes"
	RespSizeBytes = "http_response_size_bytes"
)

const (
	Method = "method"
	Path   = "path"
	Status = "status"
)
