package main

import (
	"io/ioutil"

	"k8s.io/client-go/kubernetes"

	"github.com/prometheus-operator/prometheus-operator/pkg/client/versioned"

	"github.com/penguinn/go-sdk/k8s"
	"github.com/penguinn/go-sdk/log"
)

func main() {
	// 使用config路径初始化k8s客户端
	cf, err := k8s.NewConfig(&k8s.Config{
		K8sHost:       "10.36.252.175:8443",
		K8sConfigPath: "./kube.conf",
	})
	if err != nil {
		log.Fatal(err)
	}
	k8sClient, err := kubernetes.NewForConfig(cf)
	if err != nil {
		log.Fatal(err)
	}
	if k8sClient == nil {
		log.Fatal("k8s client nil")
	}

	// 使用config字节初始化PrometheusOperator客户端
	kubeConfigBytes, err := ioutil.ReadFile("./kube.conf")
	if err != nil {
		log.Fatal(err)
	}
	cf, err = k8s.NewConfig(&k8s.Config{
		K8sHost:   "10.36.252.175:8443",
		K8sConfig: string(kubeConfigBytes),
	})
	if err != nil {
		log.Fatal(err)
	}
	prometheusOperatorClient, err := versioned.NewForConfig(cf)
	if err != nil {
		log.Fatal(err)
	}
	if prometheusOperatorClient == nil {
		log.Fatal("prometheusOperator client nil")
	}

	// 使用jwt文件初始化PrometheusOperator客户端
	cf, err = k8s.NewConfig(&k8s.Config{
		K8sHost: "10.36.252.175:8443",
		JwtPath: "./jwt.conf",
	})
	if err != nil {
		log.Fatal(err)
	}
	prometheusOperatorClient, err = versioned.NewForConfig(cf)
	if err != nil {
		log.Fatal(err)
	}
	if prometheusOperatorClient == nil {
		log.Fatal("prometheusOperator client nil")
	}

	// 使用jwt字节初始化k8s客户端
	jwtBytes, err := ioutil.ReadFile("./jwt.conf")
	if err != nil {
		log.Fatal(err)
	}
	cf, err = k8s.NewConfig(&k8s.Config{
		K8sHost: "10.36.252.175:8443",
		Jwt:     string(jwtBytes),
	})
	if err != nil {
		log.Fatal(err)
	}
	k8sClient, err = kubernetes.NewForConfig(cf)
	if err != nil {
		log.Fatal(err)
	}
	if k8sClient == nil {
		log.Fatal("k8s client nil")
	}
}
