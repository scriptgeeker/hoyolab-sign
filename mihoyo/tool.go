package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// PrintError 打印错误信息
func PrintError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

// PurifyString 删除字符串中多余空行和注释
func PurifyString(content string) string {
	list := make([]string, 0, 10)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if len(line) > 0 && line[0] != '#' {
			list = append(list, line)
		}
	}
	return strings.Join(list, "\n")
}

// GetUserInput 获取用户输入
func GetUserInput(tip string) string {
	fmt.Printf("%s", tip)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	PrintError(err)
	input = strings.TrimSpace(input)
	return input
}

// PrintAndLog 打印并记录日志
func PrintAndLog(path string, info []string) {
	line := time.Now().String()[:19] + "\n"
	for _, s := range info {
		line += "\t" + s + "\n"
	}
	line += "\n"
	fmt.Print(line)
	path = GetAbsPath(path)
	AppendFileContent(path, line)
}
