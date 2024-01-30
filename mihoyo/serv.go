package main

import (
	"fmt"
	"os"
	"strconv"
)

// SelectLang 选择语言
func SelectLang() string {
	for index, item := range LangZoneList {
		fmt.Printf("%d %s\n", index, item[0])
	}
	index := -1
	limit := len(LangZoneList) - 1
	for index < 0 || limit < index {
		tip := fmt.Sprintf(TipMap[Zone][0], 0, limit)
		num, err := strconv.Atoi(GetUserInput(tip))
		if err != nil {
			index = -1
		} else {
			index = num
		}
	}
	return LangZoneList[index][0]
}

// LoadLangZone 加载语言选项
func LoadLangZone() (string, string) {
	path := GetAbsPath(LANG_PATH)
	if !FileExists(path) {
		key := SelectLang()
		content := "# " + key + "\n" + LangZoneMap[key]
		WriteFileContent(path, content)
	}
	content := ReadFileContent(path)
	zone := PurifyString(content)
	lang, ok := ZoneLangMap[zone]
	if !ok {
		err := os.Remove(path)
		PrintError(err)
		return LoadLangZone()
	}
	return lang, zone
}

// InputCookie 获取 Cookie
func InputCookie() string {
	tip := TipMap[Zone][1]
	input := GetUserInput(tip)
	return input
}

// LoadCookie 加载 Cookie
func LoadCookie() string {
	path := GetAbsPath(COOKIE_PATH)
	if !FileExists(path) {
		input := InputCookie()
		content := "# Cookie-Editor: Export As Header String\n" + input
		WriteFileContent(path, content)
	}
	content := ReadFileContent(path)
	cookie := PurifyString(content)
	info := GetUserInfo(cookie)
	if info == nil {
		fmt.Println(TipMap[Zone][2])
		err := os.Remove(path)
		PrintError(err)
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
	info := make(map[string]string)
	temp := resp.Data.(map[string]interface{})["user_info"]
	for key, val := range temp.(map[string]interface{}) {
		info[key] = fmt.Sprint(val)
	}
	return info
}

// InputActIds 输入角色 ID
func InputActIds() string {
	content := ""
	fmt.Println(TipMap[Zone][3])
	for _, name := range GameNameMap[Zone] {
		sign := SignMap[name]
		tip := fmt.Sprintf(TipMap[Zone][4], name)
		input := GetUserInput(tip)
		if 0 < len(input) {
			content += "# " + name + "\n"
			content += sign + "=" + input + "\n"
			content += "\n"
		}
	}
	return content
}

// LoadActIds 加载角色 ID
func LoadActIds() map[string]string {
	path := GetAbsPath(ACT_IDS_PATH)
	if !FileExists(path) {
		content := InputActIds()
		WriteFileContent(path, content)
	}
	content := ReadFileContent(path)
	confStr := PurifyString(content)
	confMap := ConfStrToMap(confStr)
	if len(confMap) == 0 {
		err := os.Remove(path)
		PrintError(err)
		return LoadActIds()
	}
	return confMap
}

// HoyoLabSign 米哈游社区签到
func HoyoLabSign(cookie string, actIds map[string]string, userInfo map[string]string) {
	for _, name := range GameNameMap[Zone] {
		sign := SignMap[name]
		url := ApiMap[sign]
		id, ok := actIds[sign]
		if ok && 0 < len(id) {
			url = fmt.Sprintf("%s?lang=%s", url, Zone)
			jsonBody := fmt.Sprintf(`{"act_id":"%s","lang":"%s"}`, id, Zone)
			jsonStr := SendPostRequest(url, cookie, jsonBody)
			resp := GetResponse(jsonStr)
			email := userInfo["email"]
			PrintAndLog(SIGN_LOG_PATH, []string{
				name, email, resp.Message, url, jsonBody, jsonStr,
			})
		}
	}
}
