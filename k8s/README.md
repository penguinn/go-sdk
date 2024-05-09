# 背景
目前Kubernetes仓库包含的主要函数有：获取k8s等客户端的初始化配置

# 主要函数
## k8s客户端初始化
1. 支持五种初始化方式：K8sConfigPath、K8sConfig、JWT、JWTPath、Serviceaccount。
```go
func NewConfig(config *K8sConfig) (*rest.Config, error)
// 1. K8sConfigPath最高优先级
// 2. k8sConfig第2优先级
// 3. JwtPath第3优先级
// 4. Jwt第4优先级
// 5. ServiceAccount第5优先级
```

2. 获取的rest.Config可以用来初始化k8sClient、PrometheusOperatorClient等

## 示例
[示例参考](https://console.cloud.baidu-int.com/devops/icode/repos/baidu/det-drd/det-go-sdk/tree/master:example/k8s_example)

