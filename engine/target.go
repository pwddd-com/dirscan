package engine

import (
	"bufio"
	"dirscan/config"
	"dirscan/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var Targets []string

func InitTargets() {
	// 判断获取参数中的target是否为域名、IP地址
	// 不是域名或者IP地址，判断是否为文件、文件夹
	// 判断是否带有协议
	// 没有协议则根据配置拼接上协议
	// 最终将结果保存到Target中
	configTarget := config.Config.Scan.Target

	// 遍历数组
	for _, s := range configTarget {
		if !isValidUrl(s) {
			// 判断是否为有效文件或者文件夹
			if isValidFolderOrFile(s) {
				// 获取文件列表
				fileList := getFileList(s)
				for _, filePath := range fileList {
					// 获取文件行，判断是否为有效url地址，有效则加入
					file, err := os.Open(filePath)
					if err != nil {
						utils.PError("读取目标文件 - " + filePath + "- 失败。")
					}
					defer file.Close()
					scanner := bufio.NewScanner(file)
					for scanner.Scan() {
						line := scanner.Text()
						// 判断读取的地址是否为有效地址
						if isValidUrl(line) {
							Targets = append(Targets, line)
						} else {
							utils.PWarn("当前文件 - " + filePath + " - 中，目标 - " + s + " - 非有效地址，已忽略。")
						}
					}
					if err := scanner.Err(); err != nil {
						utils.PError("读取目标文件 - " + filePath + "- 失败。")
					}
				}
			} else {
				utils.PWarn("当前目标 - " + s + " - 非有效地址，已忽略。")
			}
		} else {
			// 添加到url地址表中
			Targets = append(Targets, s)
		}
	}

	// Target去重
	result := make(map[string]bool)
	var uniqRes []string
	for _, target := range Targets {
		if !result[target] {
			uniqRes = append(uniqRes, target)
			result[target] = true
		}
	}

	// 协议处理
	var allProtocol bool
	protocol := config.Config.Scan.Protocol
	// 是否是全部协议
	for _, s := range protocol {
		if s == "*" {
			allProtocol = true
			break
		}
	}

	// 获取uniqRes中不包含协议的
	var hasProtocolUrls []string
	for _, url := range uniqRes {
		if !strings.HasPrefix(url, "http://") {
			if !strings.HasPrefix(url, "https://") {
				if allProtocol {
					hasProtocolUrls = append(hasProtocolUrls, "http://"+url)
					hasProtocolUrls = append(hasProtocolUrls, "https://"+url)
				} else {
					for _, pro := range protocol {
						hasProtocolUrls = append(hasProtocolUrls, pro+"://"+url)
					}
				}
			} else {
				hasProtocolUrls = append(hasProtocolUrls, url)
			}
		} else {
			hasProtocolUrls = append(hasProtocolUrls, url)
		}
	}
	Targets = hasProtocolUrls
}

func getFileList(path string) []string {
	var fileList []string
	file, _ := os.Stat(path)
	mode := file.Mode()
	if mode.IsDir() {
		// 是文件夹，递归获取文件夹下的全部文件
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() { // 如果不是文件夹,是文件,存储文件名
				fileList = append(fileList, path)
			}
			return nil
		})
	} else {
		// 是文件，添加到文件列表
		fileList = append(fileList, path)
	}

	return fileList
}

func isValidFolderOrFile(s string) bool {
	_, err := os.Stat(s)
	if err != nil {
		return false
	}
	return true
}

func isValidUrl(s string) bool {
	// 是否为IPv4地址
	ipv4Matched, _ := regexp.MatchString(`(?:(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])`, s)
	if ipv4Matched {
		return ipv4Matched
	} else {
		// 是否为域名
		domainMatched, _ := regexp.MatchString(`(?:(?:[a-zZ-Z0-9]+)\.){2}((com|org|net)\.)?(com|cn|net|org|biz|info|cc|tv|top|vip)`, s)
		if domainMatched {
			return domainMatched
		} else {
			return false
		}
	}
}
