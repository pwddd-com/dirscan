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
	for _, target := range engine.Targets {
		utils.PDebug("扫描目标：" + target)
	}

	// 获取字典

}
