package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// SelectLang 选择语言
func SelectLang() LangVO {
	for index, vo := range LangList {
		fmt.Printf("%d %s\n", index, vo.Name)
	}
	index := -1
	limit := len(LangList) - 1
	for index < 0 || limit < index {
		tip := fmt.Sprintf(TipMap[Lang][0], 0, limit)
		num, err := strconv.Atoi(GetUserInput(tip))
		if err != nil {
			index = -1
		} else {
			index = num
		}
	}
	return LangList[index]
}

// LoadLang 加载语言选项
func LoadLang() string {
	path := GetAbsPath(LANG_PATH)
	if !FileExists(path) {
		vo := SelectLang()
		content := "# " + vo.Name + "\n" + vo.Lang
		WriteFileContent(path, content)
	}
	content := ReadFileContent(path)
	lang := PurifyString(content)
	_, ok := GamesMap[lang]
	if !ok {
		RemoveFile(path)
		return LoadLang()
	}
	return lang
}

// InputCookie 获取 Cookie
func InputCookie() string {
	tip := TipMap[Lang][1]
	input := GetUserInput(tip)
	return input
}

// LoadCookie 加载 Cookie
func LoadCookie() string {
	path := GetAbsPath(COOKIE_PATH)
	if !FileExists(path) {
		input := InputCookie()
		content := "# By Cookie-Editor: Export As Header String\n" + input
		WriteFileContent(path, content)
	}
	content := ReadFileContent(path)
	cookie := PurifyString(content)
	info := GetUserInfo(cookie)
	if info == nil {
		fmt.Println(TipMap[Lang][2])
		RemoveFile(path)
		return LoadCookie()
	}
	return cookie
}

// GetUserInfo 获取登录用户信息
func GetUserInfo(cookie string) map[string]string {
	url := ApiMap["TOKEN"]
	html := SendPostRequest(url, cookie, "{}")
	resp := GetResponse(html)
	if resp.Retcode != 0 {
		return nil
	}
	user := make(map[string]string)
	info := resp.Data.(map[string]interface{})["user_info"]
	for key, val := range info.(map[string]interface{}) {
		user[key] = fmt.Sprint(val)
	}
	return user
}

// DefaultActId 默认 act_id
func DefaultActId() string {
	content := ""
	for _, game := range GamesMap[Lang] {
		content += "# " + game.Name + "\n"
		content += game.ActId + "\n"
		content += "\n"
	}
	return content
}

// ParseActId 解析 act_id
func ParseActId(content string) []GameVO {
	dict := make(map[string]GameVO)
	for _, vo := range GamesMap[Lang] {
		dict[vo.Name] = vo
	}
	list := make([]GameVO, 0, 10)
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "#") && (i+1) < len(lines) {
			line = strings.TrimPrefix(line, "#")
			id := strings.TrimSpace(lines[i+1])
			vo, ok := dict[strings.TrimSpace(line)]
			if ok {
				vo.ActId = id
				list = append(list, vo)
			}

		}
	}
	return list
}

// UpdateGamesMap 更新游戏信息
func UpdateGamesMap() {
	path := GetAbsPath(ACT_IDS_PATH)
	if !FileExists(path) {
		content := DefaultActId()
		WriteFileContent(path, content)
	}
	content := ReadFileContent(path)
	GamesMap[Lang] = ParseActId(content)
}

// HoyoLabSign 米哈游社区签到
func HoyoLabSign(userInfo map[string]string) {
	for _, game := range GamesMap[Lang] {
		url := fmt.Sprintf("%s?lang=%s", game.SignUrl, Lang)
		jsonBody := fmt.Sprintf(`{"act_id":"%s","lang":"%s"}`, game.ActId, Lang)
		jsonResp := SendPostRequest(url, Cookie, jsonBody)
		msg := GetResponse(jsonResp).Message
		if msg == "OK" {
			msg = TipMap[Lang][3]
		}
		email := userInfo["email"]
		PrintAndLog(SIGN_LOG_PATH, []string{
			game.Name, email, msg, url, jsonBody, jsonResp,
		})
	}
}

// TimedSign 定时签到
func TimedSign() {
	// 恐慌异常处理
	defer func() {
		err := recover()
		if err != nil {
			PrintError(errors.New(fmt.Sprint(err)))
			time.Sleep(60 * time.Second) // 1 分钟后重试
			TimedSign()
		}
	}()
	// 开启定时任务
	SetInterval(Interval, func() error {
		userInfo := GetUserInfo(Cookie)
		if userInfo == nil {
			return errors.New("cookie has expired")
		}
		HoyoLabSign(userInfo)
		return nil
	})
}
