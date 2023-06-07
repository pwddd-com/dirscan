package main

import (
	"dirscan/cmd"
	"dirscan/config"
	"dirscan/engine"
	"dirscan/utils"
	"fmt"
	"time"
)

func main() {
	banner := `
 ██████████   █████ ███████████    █████████    █████████    █████████   ██████   █████
░░███░░░░███ ░░███ ░░███░░░░░███  ███░░░░░███  ███░░░░░███  ███░░░░░███ ░░██████ ░░███ 
 ░███   ░░███ ░███  ░███    ░███ ░███    ░░░  ███     ░░░  ░███    ░███  ░███░███ ░███ 
 ░███    ░███ ░███  ░██████████  ░░█████████ ░███          ░███████████  ░███░░███░███ 
 ░███    ░███ ░███  ░███░░░░░███  ░░░░░░░░███░███          ░███░░░░░███  ░███ ░░██████ 
 ░███    ███  ░███  ░███    ░███  ███    ░███░░███     ███ ░███    ░███  ░███  ░░█████ 
 ██████████   █████ █████   █████░░█████████  ░░█████████  █████   █████ █████  ░░█████
░░░░░░░░░░   ░░░░░ ░░░░░   ░░░░░  ░░░░░░░░░    ░░░░░░░░░  ░░░░░   ░░░░░ ░░░░░    ░░░░░ 

[+] code by Mufeng V0.0.1
[+] https://github.com/marmufeng/Dirscan
`
	fmt.Println(banner)

	startTime := time.Now()
	utils.PInfo("当前系统时间：" + startTime.String())

	// 初始化配置信息
	config.InitConfig()

	// 加载命令行
	cmd.Execute()

	// 扫描
	// 获取目标
	engine.InitTargets()
	for _, target := range engine.Targets {
		fmt.Println(target)
	}

	// 获取字典

}
