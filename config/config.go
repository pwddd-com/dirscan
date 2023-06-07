package config

import (
	"dirscan/utils"
	"github.com/spf13/viper"
	"os"
)

// ScanConfiguration 扫描配置
type ScanConfiguration struct {
	Target   []string `yaml:"target"`
	Protocol []string `yaml:"protocol"`
	Language []string `yaml:"language"`
	Dict     []string `yaml:"dict"`
}

// ViewConfiguration 显示设置
type ViewConfiguration struct {
	ConsoleLog bool   `yaml:"consoleLog"`
	OutputType string `yaml:"outputType"`
	OutputFile string `yaml:"outputFile"`
	ResultPush bool   `yaml:"resultPush"`
	PushUrl    string `yaml:"pushUrl"`
	Out2File   bool   `yaml:"out2File"`
}

// HeaderConfiguration http设置
type HeaderConfiguration struct {
	UserAgent []string `yaml:"UserAgent"`
	Cookie    string   `yaml:"cookie"`
	Others    []string `yaml:"others"`
}

type RequestConfiguration struct {
	Timeout      int      `yaml:"timeout"`
	FailRetries  int      `yaml:"failRetries"`
	IgnoreStatus []string `yaml:"ignoreStatus"`
	Concurrency  int      `yaml:"concurrency"`
}

type HttpConfiguration struct {
	Http2   bool                 `yaml:"http2"`
	Proxy   string               `yaml:"proxy"`
	Headers HeaderConfiguration  `yaml:"headers"`
	Request RequestConfiguration `yaml:"request"`
}

type Configuration struct {
	Version string            `yaml:"version"`
	Enabled bool              `yaml:"enabled"`
	Scan    ScanConfiguration `yaml:"scan"`
	View    ViewConfiguration `yaml:"view"`
	Http    HttpConfiguration `yaml:"http"`
}

func (configuration *Configuration) getConfiguration() *Configuration {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	vip := viper.New()
	vip.AddConfigPath(path)
	vip.SetConfigName("config")
	vip.SetConfigType("yaml")

	if err := vip.ReadInConfig(); err != nil {
		utils.PError("配置读取失败，请检测可执行文件同级目录下是否存在config.yaml")
	}

	err = vip.Unmarshal(&configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}

var Config *Configuration

func InitConfig() {
	utils.PInfo("开始读取应用配置文件：config.yaml")
	Config = new(Configuration)
	Config.getConfiguration()
	if Config.Enabled {
		utils.PInfo("配置加载完成。")
	} else {
		Config = nil
		utils.PWarn("当前系统配置文件为非启用状态，启用请编辑配置文件设置 Enabled 为 true ")
	}
}

// SysConfig 系统默认配置
func SysConfig() {
	// 默认使用 * 协议
	Config.Scan.Protocol = append(Config.Scan.Protocol, "*")

	// 默认使用Java语言
	Config.Scan.Language = append(Config.Scan.Language, "java")

	// 默认开启日志打印，关闭结果推送和输出到文件
	Config.View.ConsoleLog = true
	Config.View.ResultPush = false
	Config.View.Out2File = false

	// 默认关闭Http2 不使用代理 不设置请求头和Cookie
	Config.Http.Http2 = false
	Config.Http.Proxy = ""
	// 默认失败不重试 超时时间5s，并发100
	Config.Http.Request.FailRetries = 0
	Config.Http.Request.Timeout = 5
	Config.Http.Request.Concurrency = 100
}

func NecessaryParam() {
	// 判断必要参数是否为空或者不存在
	// 必要参数包括 扫描目标 字典
	if len(Config.Scan.Target) == 0 || len(Config.Scan.Dict) == 0 {
		utils.PError("应用程序缺少必要的参数（扫描目标或者字典），将无法启动，请检查是否指定参数。")
		os.Exit(0)
	}
}
