package main

import (
	"github.com/penguinn/go-sdk/example/metrics_example/constant"
	"github.com/penguinn/go-sdk/log"
	"github.com/penguinn/go-sdk/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

// 申明Counter/Gauge/Histogram
var metricsValues = []*metrics.MetricValue{
	{
		ValueType: metrics.Counter,                                                                  // 提供四种指标类型Counter,Gauge,Histogram,Summary
		Name:      constant.MetricTranslateLanguage,                                                 // 本服务唯一
		Help:      "hello interface count",                                                          // 指标描述
		Labels:    []string{constant.MetricLabelSourceLanguage, constant.MetricLabelTargetLanguage}, // 标签注意顺序
	},
	{
		ValueType: metrics.Histogram,
		Name:      constant.MetricRandomNumber,
		Help:      "A histogram of normally distributed random numbers.",
		Buckets:   prometheus.LinearBuckets(-3, .1, 61),
	},
}

func HistogramInit() {
	err := metrics.InitMetric(metricsValues)
	if err != nil {
		log.Fatal(err)
	}
}
