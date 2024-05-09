package main

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/penguinn/go-sdk/example/metrics_example/api/handler"
	"github.com/penguinn/go-sdk/log"
	"github.com/penguinn/go-sdk/metrics"
)

func main() {
	g := gin.Default()

	// metrics_example sdk需要初始化的
	_ = metrics.Init(g, "bml", "modelRepo")
	CounterGuageInit()
	HistogramInit()

	// gin需要初始化的
	handler.HelloHandler{}.Init(g)
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalln(err)
	}
	err = g.RunListener(listener)
	if err != nil {
		log.Fatal(err)
	}
}
