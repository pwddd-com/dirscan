package utils

import (
	"os"
	"path/filepath"
)

func IsValidFolderOrFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func GetFileList(path string) []string {
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
