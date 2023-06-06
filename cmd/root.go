package cmd

import (
	"dirscan/config"
	"github.com/spf13/cobra"
	"strings"
)

// 根命令
var rootCmd = &cobra.Command{
	Use:   "dirscan",
	Short: "An Dir Scan CLI",
	Long:  `This is an Dir Scan CLI for web security.`,
	Run: func(cmd *cobra.Command, args []string) {
		initConfigValues(cmd)

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
	viewConfig := config.Config.View
	rootCmd.Flags().StringP("output", "o", viewConfig.OutputFile+"."+config.Config.View.OutputType, "Output result as a file.")
	rootCmd.Flags().BoolP("view", "v", viewConfig.ConsoleLog, "Print console log.")
	rootCmd.Flags().StringP("push", "p", viewConfig.PushUrl, "Push result to webserver.")
}

func initViewValues(cmd *cobra.Command) {
	viewConfig := config.Config.View
	output, _ := cmd.Flags().GetString("output")
	if output != viewConfig.OutputFile+"."+viewConfig.OutputType {
		if strings.LastIndex(output, ".") != -1 {
			viewConfig.OutputFile = output[0:strings.LastIndex(output, ".")]
			viewConfig.OutputType = output[strings.LastIndex(output, ".")+1:]
		} else {
			viewConfig.OutputFile = output
		}
	}

	viewConfig.ConsoleLog, _ = cmd.Flags().GetBool("view")

	pushUrl, _ := cmd.Flags().GetString("push")
	if pushUrl != viewConfig.PushUrl {
		viewConfig.PushUrl = pushUrl
		viewConfig.ResultPush = true
	}

}

func initHttpFlags() {
	httpConfig := config.Config.Http
	rootCmd.Flags().StringP("proxy", "", httpConfig.Proxy, "Add a Proxy.")
	rootCmd.Flags().StringArrayP("user-agent", "u", httpConfig.Headers.UserAgent, "Specify User-Agents.")
	rootCmd.Flags().StringP("cookie", "c", httpConfig.Headers.Cookie, "Specify Cookie.")
	rootCmd.Flags().StringArrayP("ignore-status", "i", httpConfig.Request.IgnoreStatus, "Ignore result status code.")
	rootCmd.Flags().StringArrayP("header", "", httpConfig.Headers.Others, "Add a Header.")
	rootCmd.Flags().IntP("qps", "q", httpConfig.Request.MaxQps, "Max requests one second.")
}

func initHttpValues(cmd *cobra.Command) {
	httpConfig := config.Config.Http
	httpConfig.Proxy, _ = cmd.Flags().GetString("proxy")
	httpConfig.Headers.UserAgent, _ = cmd.Flags().GetStringArray("user-agent")
	httpConfig.Headers.Cookie, _ = cmd.Flags().GetString("cookie")
	httpConfig.Request.IgnoreStatus, _ = cmd.Flags().GetStringArray("ignore-status")
	httpConfig.Headers.Others, _ = cmd.Flags().GetStringArray("header")
	httpConfig.Request.MaxQps, _ = cmd.Flags().GetInt("qps")
}

func initScanFlags() {
	scanConfig := config.Config.Scan
	rootCmd.Flags().StringArrayP("language", "l", scanConfig.Language, "Target language.")
	rootCmd.Flags().StringArrayP("target", "t", scanConfig.Target, "Specify targets.")
	rootCmd.Flags().StringArrayP("protocol", "", scanConfig.Protocol, "Specify protocol.")
	rootCmd.Flags().StringArrayP("dict", "d", scanConfig.Dict, "Specify a dict.")
}

func initScanValues(cmd *cobra.Command) {
	scanConfig := config.Config.Scan
	scanConfig.Language, _ = cmd.Flags().GetStringArray("language")
	config.Config.Scan.Dict, _ = cmd.Flags().GetStringArray("dict")
	config.Config.Scan.Protocol, _ = cmd.Flags().GetStringArray("protocol")
	config.Config.Scan.Target, _ = cmd.Flags().GetStringArray("target")
}

func Execute() {
	rootCmd.Execute()
}
