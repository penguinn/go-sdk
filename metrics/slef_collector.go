package metrics

import (
	"sort"

	"github.com/golang/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

type Collector interface {
	Describe(chan<- *prometheus.Desc)
	Collect(chan<- prometheus.Metric)
}

type SelfCollector struct {
	valueFunc
	desc      *prometheus.Desc
	valueType prometheus.ValueType
	function  func() []*LabelAndValue
}

func NewSelfCollector(desc *prometheus.Desc, valueType prometheus.ValueType, function func() []*LabelAndValue) *SelfCollector {
	return &SelfCollector{
		desc:      desc,
		valueType: valueType,
		function:  function,
	}
}

func (c *SelfCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.desc
}

func (c *SelfCollector) Collect(ch chan<- prometheus.Metric) {
	for _, labelAndValue := range c.function() {
		labelPairs := make([]*dto.LabelPair, 0, len(labelAndValue.Labels))
		for name, value := range labelAndValue.Labels {
			labelPairs = append(labelPairs, &dto.LabelPair{
				Name:  proto.String(name),
				Value: proto.String(value),
			})
		}
		sort.Sort(labelPairSorter(labelPairs))

		v := &valueFunc{
			desc:       c.desc,
			valType:    c.valueType,
			value:      labelAndValue.Value,
			labelPairs: labelPairs,
		}

		ch <- v
	}
}

type labelPairSorter []*dto.LabelPair

func (s labelPairSorter) Len() int {
	return len(s)
}

func (s labelPairSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s labelPairSorter) Less(i, j int) bool {
	return s[i].GetName() < s[j].GetName()
}
