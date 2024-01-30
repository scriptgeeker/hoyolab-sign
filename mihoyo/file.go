package main

import (
	"os"
	"path/filepath"
)

// GetAbsPath 获取绝对路径
func GetAbsPath(path string) string {
	absPath, _ := filepath.Abs(path)
	return absPath
}

// FileExists 判断文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ReadFileContent 读取文件内容
func ReadFileContent(path string) string {
	content, _ := os.ReadFile(path)
	return string(content)
}

// WriteFileContent 写入文件内容
func WriteFileContent(path string, content string) {
	err := os.WriteFile(path, []byte(content), 666)
	PrintError(err)
}

// AppendFileContent 追加文件内容
func AppendFileContent(path string, content string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	PrintError(err)
	defer func(file *os.File) {
		err := file.Close()
		PrintError(err)
	}(file)
	_, err = file.Write([]byte(content))
	PrintError(err)
}
