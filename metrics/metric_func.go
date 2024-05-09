package metrics

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"

	"github.com/penguinn/go-sdk/log"
)

type LabelAndValue struct {
	Labels map[string]string
	Value  float64
}

func NewCounterFunc(opts prometheus.CounterOpts, function func() []*LabelAndValue) prometheus.CounterFunc {
	counterFunc := newCounterFunc(opts, function)
	prometheus.DefaultRegisterer.MustRegister(counterFunc)
	return counterFunc
}

func newCounterFunc(opts prometheus.CounterOpts, function func() []*LabelAndValue) prometheus.CounterFunc {
	return NewSelfCollector(prometheus.NewDesc(
		prometheus.BuildFQName(opts.Namespace, opts.Subsystem, opts.Name),
		opts.Help,
		nil,
		opts.ConstLabels,
	), prometheus.CounterValue, function)
}

func NewGaugeFunc(opts prometheus.GaugeOpts, function func() []*LabelAndValue) prometheus.GaugeFunc {
	gaugeFunc := newGaugeFunc(opts, function)
	prometheus.DefaultRegisterer.MustRegister(gaugeFunc)
	return gaugeFunc
}

func newGaugeFunc(opts prometheus.GaugeOpts, function func() []*LabelAndValue) prometheus.GaugeFunc {
	return NewSelfCollector(prometheus.NewDesc(
		prometheus.BuildFQName(opts.Namespace, opts.Subsystem, opts.Name),
		opts.Help,
		nil,
		opts.ConstLabels,
	), prometheus.GaugeValue, function)
}

type valueFunc struct {
	desc       *prometheus.Desc
	valType    prometheus.ValueType
	value      float64
	labelPairs []*dto.LabelPair
}

func (v *valueFunc) Desc() *prometheus.Desc {
	return v.desc
}

func (v *valueFunc) Write(out *dto.Metric) error {
	err := populateMetric(v.valType, v.value, v.labelPairs, nil, out)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func populateMetric(t prometheus.ValueType, v float64, labelPairs []*dto.LabelPair, e *dto.Exemplar, m *dto.Metric) error {
	m.Label = labelPairs
	switch t {
	case prometheus.CounterValue:
		m.Counter = &dto.Counter{Value: proto.Float64(v), Exemplar: e}
	case prometheus.GaugeValue:
		m.Gauge = &dto.Gauge{Value: proto.Float64(v)}
	case prometheus.UntypedValue:
		m.Untyped = &dto.Untyped{Value: proto.Float64(v)}
	default:
		return fmt.Errorf("encountered unknown type %v", t)
	}
	return nil
}
