package main

import (
	"dirscan/cmd"
	"dirscan/config"
	"dirscan/engine"
	"dirscan/utils"
	"strconv"
	"time"
)

func main() {

	utils.PBanner()

	startTime := time.Now()
	utils.PInfo("当前系统时间：" + startTime.Format("2006-01-02 15:04:05"))

	// 初始化配置信息
	config.InitConfig()
	utils.PDebug("当前配置文件是否为空：" + strconv.FormatBool(config.Config == nil))

	// 加载命令行
	cmd.Execute()

	// 必要参数判断
	config.NecessaryParam()

	// 扫描
	// 获取目标
	utils.PInfo("即将开始处理扫描目标。")
	initTargetTime := time.Now()
	engine.InitTargets()
	utils.PInfo("扫描目标处理完成，扫描目标数量：" + strconv.Itoa(len(engine.Targets)) + " 个，耗时：" + strconv.FormatFloat(time.Now().Sub(initTargetTime).Seconds(), 'f', 2, 64) + " s")

	// 获取字典
	utils.PInfo("即将开始收集和处理扫描字典。")
	initDictTime := time.Now()
	engine.InitDict()
	utils.PInfo("扫描字典处理完成，扫描字典数量：" + strconv.Itoa(len(engine.Dict)) + " 个，耗时：" + strconv.FormatFloat(time.Now().Sub(initDictTime).Seconds(), 'f', 2, 64) + " s")

	// 计算扫描次数
	totalRequestCount := len(engine.Dict) * len(engine.Targets)
	scanStartTime := time.Now()
	utils.PInfo("前置处理完成，准备开始扫描。总请求数量为： " + strconv.Itoa(totalRequestCount) + " 个。当前时间： " + scanStartTime.Format("2006-01-02 15:04:05"))

	// 开始扫描啦

}
