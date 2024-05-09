package main

import (
	"github.com/penguinn/go-sdk/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

func CounterGuageInit() {
	metrics.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "bml",
		Subsystem: "notebook",
		Name:      "name1",
		Help:      "here help name1",
	}, TestGuage)

	metrics.NewCounterFunc(prometheus.CounterOpts{
		Namespace: "bml",
		Subsystem: "notebook",
		Name:      "name2",
		Help:      "here help name2",
	}, TestCounter)
}

func TestGuage() []*metrics.LabelAndValue {
	labelAndValues := []*metrics.LabelAndValue{
		{Labels: map[string]string{"project": "test", "organization": "bml"}, Value: 5},
		{Labels: map[string]string{"project": "test", "organization": "easydata"}, Value: 6},
	}
	return labelAndValues
}

func TestCounter() []*metrics.LabelAndValue {
	labelAndValues := []*metrics.LabelAndValue{{Labels: map[string]string{"organization": "test1"}, Value: 10}}
	return labelAndValues
}
