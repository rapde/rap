package utils

import (
	"log"
	"os"
	"os/exec"
)

// GetWorkDir 获取运行目录
func GetWorkDir() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get work dir %v", err)
	}

	return wd
}

// GetBinPath 从当前环境中获取 可执行文件 路径
func GetBinPath(binName string) (path string, exist bool) {
	path, err := exec.LookPath(binName)
	if err != nil {
		return "", false
	}

	return path, true
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.Mode().IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.Mode().IsRegular()
}
