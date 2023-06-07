package cmd

import (
	"dirscan/config"
	"dirscan/utils"
	"github.com/spf13/cobra"
	"strings"
)

// 根命令
var rootCmd = &cobra.Command{
	Use:   "dirscan",
	Short: "An Dir Scan CLI",
	Long:  `This is an Dir Scan CLI for web security.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PInfo("开始解析命令行参数，命令行参数将会覆盖配置文件配置。")
		if config.Config == nil {
			config.Config = new(config.Configuration)
			// 使用系统默认配置
			config.SysConfig()
		}
		initConfigValues(cmd)
		utils.PInfo("命令行参数解析完成。")
	},
}

func init() {
	initFlags()
}

// 从命令行获取参数
func initFlags() {
	// 获取扫描对象设置
	initScanFlags()
	// 获取显示设置
	initViewFlags()
	// 获取请求设置
	initHttpFlags()
}

func initConfigValues(cmd *cobra.Command) {
	initScanValues(cmd)
	initHttpValues(cmd)
	initViewValues(cmd)
}

func initViewFlags() {
	rootCmd.Flags().StringP("output", "o", "", "Output result as a file.")
	rootCmd.Flags().BoolP("view", "v", false, "Print console log.")
	rootCmd.Flags().StringP("push", "p", "", "Push result to webserver.")
}

func initViewValues(cmd *cobra.Command) {
	viewConfig := config.Config.View
	output, _ := cmd.Flags().GetString("output")
	if output != "" {
		viewConfig.Out2File = true
		if strings.LastIndex(output, ".") != -1 {
			viewConfig.OutputFile = output[0:strings.LastIndex(output, ".")]
			viewConfig.OutputType = output[strings.LastIndex(output, ".")+1:]
		} else {
			viewConfig.OutputFile = output
		}
	}

	consoleLog, _ := cmd.Flags().GetBool("view")
	if consoleLog {
		viewConfig.ConsoleLog = consoleLog
	}

	pushUrl, _ := cmd.Flags().GetString("push")
	if pushUrl != "" {
		viewConfig.PushUrl = pushUrl
		viewConfig.ResultPush = true
	}
}

func initHttpFlags() {
	rootCmd.Flags().StringP("proxy", "", "", "Add a Proxy.")
	rootCmd.Flags().StringArrayP("user-agent", "u", nil, "Specify User-Agents.")
	rootCmd.Flags().StringP("cookie", "c", "", "Specify Cookie.")
	rootCmd.Flags().StringArrayP("ignore-status", "i", nil, "Ignore result status code.")
	rootCmd.Flags().StringArrayP("header", "", nil, "Add a Header.")
	rootCmd.Flags().IntP("qps", "q", 0, "Max requests one second.")
}

func initHttpValues(cmd *cobra.Command) {
	httpConfig := config.Config.Http
	proxy, _ := cmd.Flags().GetString("proxy")
	userAgent, _ := cmd.Flags().GetStringArray("user-agent")
	cookie, _ := cmd.Flags().GetString("cookie")
	ignoreStatus, _ := cmd.Flags().GetStringArray("ignore-status")
	headers, _ := cmd.Flags().GetStringArray("header")
	maxQps, _ := cmd.Flags().GetInt("qps")

	if proxy != "" {
		httpConfig.Proxy = proxy
	}

	if userAgent != nil {
		httpConfig.Headers.UserAgent = userAgent
	}

	if cookie != "" {
		httpConfig.Headers.Cookie = cookie
	}

	if ignoreStatus != nil {
		httpConfig.Request.IgnoreStatus = ignoreStatus
	}

	if headers != nil {
		httpConfig.Headers.Others = headers
	}

	if maxQps != -1 {
		httpConfig.Request.MaxQps = maxQps
	}
}

func initScanFlags() {
	rootCmd.Flags().StringArrayP("language", "l", nil, "Target language.")
	rootCmd.Flags().StringArrayP("target", "t", nil, "Specify targets.")
	rootCmd.Flags().StringArrayP("protocol", "", nil, "Specify protocol.")
	rootCmd.Flags().StringArrayP("dict", "d", nil, "Specify a dict.")
}

func initScanValues(cmd *cobra.Command) {
	scanConfig := config.Config.Scan
	language, _ := cmd.Flags().GetStringArray("language")
	dict, _ := cmd.Flags().GetStringArray("dict")
	protocol, _ := cmd.Flags().GetStringArray("protocol")
	target, _ := cmd.Flags().GetStringArray("target")

	if language != nil {
		scanConfig.Language = language
	}

	if dict != nil {
		scanConfig.Dict = dict
	}

	if protocol != nil {
		scanConfig.Protocol = protocol
	}

	if target != nil {
		scanConfig.Target = target
	}
}

func Execute() {
	rootCmd.Execute()
}
