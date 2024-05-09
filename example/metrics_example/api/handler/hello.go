package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/penguinn/go-sdk/example/metrics_example/constant"
	"github.com/penguinn/go-sdk/log"
	"github.com/penguinn/go-sdk/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type HelloHandler struct {
}

func (p *HelloHandler) Hello(c *gin.Context) {
	// 业务打点
	// metrics包维护默认的metric管理器
	// 使用IncWithLabel给带有标签的指标加1
	labels := map[string]string{constant.MetricLabelSourceLanguage: "cn", constant.MetricLabelTargetLanguage: "en"}
	err := metrics.IncWithLabel(constant.MetricTranslateLanguage, labels)
	if err != nil {
		log.Error(err)
		c.Status(500)
		return
	}

	// 使用Inc给带有标签的指标加1，不过这个方法需要让初始化的label和value的顺序一样。如下所示：
	err = metrics.Inc(constant.MetricTranslateLanguage, []string{"cn", "en"})
	if err != nil {
		log.Error(err)
		c.Status(500)
		return
	}

	log.Info("hello world")
	_, _ = c.Writer.Write([]byte("hello world"))
}

func (p *HelloHandler) World(c *gin.Context) {
	// 如果metrics库的方法不能完全满足你, 那么可以获取Vec后转换成相应的指标并使用
	metric, err := metrics.GetMetric(constant.MetricRandomNumber)
	if err != nil {
		log.Error(err)
		c.Status(500)
		return
	}
	metric.GetVec().(prometheus.Histogram).Observe(100)
}

func (p HelloHandler) Init(g *gin.Engine) {
	g.GET("/hello", p.Hello)
	g.GET("/world", p.World)
}
