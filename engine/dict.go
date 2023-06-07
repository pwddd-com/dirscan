package engine

import (
	"bufio"
	"dirscan/config"
	"dirscan/utils"
	"os"
)

var Dict []string

func InitDict() {
	// 获取配置中的dict配置
	dictList := config.Config.Scan.Dict
	for _, path := range dictList {
		// 判断是否是文件
		if utils.IsValidFolderOrFile(path) {
			// 有效文件
			fileList := utils.GetFileList(path)
			for _, p := range fileList {
				file, err := os.Open(p)
				if err != nil {
					utils.PError("读取目标文件 - " + p + "- 失败。")
				}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					line := scanner.Text()
					// 判断读取的地址是否为有效地址
					Dict = append(Dict, line)
				}
				if err := scanner.Err(); err != nil {
					utils.PError("读取目标文件 - " + p + "- 失败。")
				}
			}
		} else {
			utils.PWarn("当前目标 - " + path + " - 非有效文件，已忽略。")
		}
	}

	// 去重
	result := make(map[string]bool)
	var uniqDic []string
	for _, dic := range Dict {
		if !result[dic] {
			uniqDic = append(uniqDic, dic)
			result[dic] = true
		}
	}
	Dict = uniqDic

}
