package k8s

import (
	"errors"
	"time"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/penguinn/go-sdk/log"
)

type Config struct {
	K8sConfigPath    string // K8S的CA和token文件路径
	K8sConfig        string // K8S的CA和token值
	JwtPath          string // Bearer Token路径
	Jwt              string // Bearer Token值
	K8sHost          string // kubernetes这个service的访问地址
	K8sClientQPS     int
	K8sClientBurst   int
	K8sClientTimeout int
}

// 1. K8sConfigPath最高优先级
// 2. k8sConfig第2优先级
// 3. JwtPath第3优先级
// 4. Jwt第4优先级
// 5. ServiceAccount第5优先级
func NewConfig(config *Config) (*rest.Config, error) {
	var err error
	var cfg *rest.Config
	configMode := 0
	if config.K8sConfigPath != "" {
		configMode++
	}
	if config.K8sConfig != "" {
		configMode++
	}
	if config.JwtPath != "" {
		configMode++
	}
	if config.Jwt != "" {
		configMode++
	}
	if configMode > 1 {
		return nil, errors.New("k8s client init mode just need 1")
	}

	if config.K8sConfigPath != "" {
		log.Info("use k8sConfig path init k8s client")
		cfg, err = NewK8sClientFromConfigPath(config)
	} else if config.K8sConfig != "" {
		log.Info("use k8sConfig init k8s client")
		cfg, err = NewK8sClientFromConfig(config)
	} else if config.JwtPath != "" {
		log.Info("use JwtPath init k8s client")
		cfg, err = NewK8SFromBearerTokenFile(config)
	} else if config.Jwt != "" {
		log.Info("use Jwt init k8s client")
		cfg, err = NewK8SFromBearerToken(config)
	} else {
		log.Info("use K8sConfigPath init k8s client")
		cfg, err = NewFromServiceAccount(config)
	}
	return cfg, err
}

// NewK8sClientFromConfigPath 使用kubeconfig路径初始化
func NewK8sClientFromConfigPath(config *Config) (*rest.Config, error) {
	if config.K8sConfigPath == "" {
		log.Infof("use InClusterConfig init k8s client")
	}
	cfg, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: config.K8sConfigPath},
		&clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, err
	}
	cfg.QPS = float32(config.K8sClientQPS)
	cfg.Burst = config.K8sClientBurst
	cfg.Timeout = time.Second * time.Duration(config.K8sClientTimeout)
	cfg.Host = config.K8sHost

	return cfg, err
}

// NewK8sClientFromConfig 使用kubeconfig字符串初始化
func NewK8sClientFromConfig(config *Config) (*rest.Config, error) {
	k8sConfig, err := clientcmd.Load([]byte(config.K8sConfig))
	if err != nil {
		log.Errorf("Load k8s clientConfig from string failed, config=[%s], %v", config.K8sConfig, err)
		return nil, err
	}
	clientConfig := clientcmd.NewDefaultClientConfig(*k8sConfig, &clientcmd.ConfigOverrides{})
	cfg, err := clientConfig.ClientConfig()
	if err != nil {
		log.Errorf("Get k8s clientConfig from string failed, config=[%v], %v", k8sConfig, err)
		return nil, err
	}
	cfg.Host = config.K8sHost
	cfg.QPS = float32(config.K8sClientQPS)
	cfg.Burst = config.K8sClientBurst
	cfg.Timeout = time.Second * time.Duration(config.K8sClientTimeout)

	return cfg, nil
}

// NewK8SFromBearerToken 使用token初始化
func NewK8SFromBearerToken(config *Config) (*rest.Config, error) {
	cfg := &rest.Config{
		Host:        config.K8sHost,
		BearerToken: config.Jwt,
		QPS:         float32(config.K8sClientQPS),
		Burst:       config.K8sClientBurst,
		Timeout:     time.Second * time.Duration(config.K8sClientTimeout),
	}

	return cfg, nil
}

// NewK8SFromBearerTokenFile 使用挂载的tokenfile路径初始化
func NewK8SFromBearerTokenFile(config *Config) (*rest.Config, error) {
	cfg := &rest.Config{
		Host:            config.K8sHost,
		BearerTokenFile: config.JwtPath,
		QPS:             float32(config.K8sClientQPS),
		Burst:           config.K8sClientBurst,
		Timeout:         time.Second * time.Duration(config.K8sClientTimeout),
	}

	return cfg, nil
}

// NewFromServiceAccount 通过serviceAccount初始化
func NewFromServiceAccount(config *Config) (*rest.Config, error) {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	cfg.QPS = float32(config.K8sClientQPS)
	cfg.Burst = config.K8sClientBurst
	cfg.Timeout = time.Second * time.Duration(config.K8sClientTimeout)

	return cfg, nil
}
