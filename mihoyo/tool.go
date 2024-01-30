package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// PrintError 打印错误信息
func PrintError(err error) {
	if err != nil {
		fmt.Println(err.Error())
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

// JsonStrToMap JSON 字符串转字典对象
func JsonStrToMap(jsonStr string) map[string]interface{} {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	PrintError(err)
	return data
}

// GetResponse 获取响应信息
func GetResponse(jsonStr string) *ResponseVO {
	vo := ResponseVO{}
	err := json.Unmarshal([]byte(jsonStr), &vo)
	PrintError(err)
	return &vo
}

// ConfStrToMap CONF 字符串转字典对象
func ConfStrToMap(confStr string) map[string]string {
	confMap := make(map[string]string)
	lines := strings.Split(confStr, "\n")
	for _, line := range lines {
		if strings.IndexAny(line, "=") != -1 {
			split := strings.Split(line, "=")
			if len(split) == 2 && 0 < len(split[1]) {
				confMap[split[0]] = split[1]
			}
		}
	}
	return confMap
}

// GetRandomInt 获取随机数
func GetRandomInt(min int, max int) int {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	num := r.Intn(max - min)
	return min + num
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
