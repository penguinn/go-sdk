# 前言
本SDK是为了方便配置的初始化

# 功能
目前只支持如下配置方式，优先级从低到高：
1. 默认值
2. 配置文件(JSON, TOML, YAML, HCL, envfile 和 Java properties格式的配置文件)
3. 环境变量

# 样例
## 文件
* config.ini
```
[Server]
Port=":9090"
[UserManager]
Method="POST"
```

* 启动程序
```go
// 环境变量设置全大写，设置名称为Key+"_"+SubKey...
// 例如ServerConfig中的Port设置方法为os.SetEnv("SERVER_PORT")
type Config struct {
	Server      ServerConfig      `mapstructure:"Server"`
	UserManager UserManagerConfig `mapstructure:"UserManager"`
}

type ServerConfig struct {
	Port string `mapstructure:"Port"`
}

type UserManagerConfig struct {
	Method string `mapstructure:"Method"`
	Url    string `mapstructure:"URL"`
}

// 实例化结构体，同时设置默认值
var GlobalConfig = Config{
	Server: ServerConfig{
		Port: ":8080",
	},
    UserManager: UserManagerConfig{
		Method: "GET",
	},
}

var configPath string
var configFileType string

func main() {
	flag.StringVar(&configFileType, "t", "ini", "config file type")
	flag.StringVar(&configPath, "f", "./config.ini", "config file path")
	flag.Parse()

	err := config.Init(configFileType, configPath, &GlobalConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(viper.AllSettings())
}
```

## 运行
* 执行
```bash
export USERMANAGER_METHOD="DELETE"; go run .
```

* 输出
```
map[server:map[port::9090] usermanager:map[method:DELETE url:]]
```

## 分析
1. 默认端口 **:8080** 被配置文件的 **:9090**覆盖
2. 默认方法 **GET** 被配置文件的 **POST** 覆盖， 最后被环境变量 **DELETE** 覆盖

## 源码
[源码参考](https://console.cloud.baidu-int.com/devops/icode/repos/baidu/det-drd/det-go-sdk/tree/master:example/config_example)



