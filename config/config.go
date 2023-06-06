package config

import (
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
	MaxQps       int      `yaml:"maxQps"`
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
		panic(err)
	}

	err = vip.Unmarshal(&configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}

var Config Configuration

func init() {
	Config.getConfiguration()
}
