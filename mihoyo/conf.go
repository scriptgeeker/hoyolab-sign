package main

// COOKIE_PATH cookie 配置文件
const COOKIE_PATH = "./cookie.txt"

// LANG_PATH 语言配置文件
const LANG_PATH = "./lang.txt"

// ACT_IDS_PATH 角色id列表
const ACT_IDS_PATH = "./act-ids.txt"

// SIGN_LOG_PATH 签到日志记录
const SIGN_LOG_PATH = "./sign.log"

// LangZoneList 语言地区列表
var LangZoneList = [][]string{
	{"简体中文", "zh-cn"},
	{"繁體中文", "zh-tw"},
	{"日本語", "ja-jp"},
	{"한국어", "ko-kr"},
	{"English", "en-us"},
}

var TipMap = map[string][]string{
	"zh-cn": {
		"请选择语言 (%d-%d): ",
		"请输入 Cookie: ",
		"该 Cookie 不正确, 请重新输入: ",
		"请输入 act_id (e开头,长度16): ",
		"%s (回车跳过): ",
	},
	"zh-tw": {
		"請選擇語言 (%d-%d): ",
		"請輸入 Cookie: ",
		"該 Cookie 不正確, 請重新輸入: ",
		"請輸入 act_id (e開頭，長度16): ",
		"%s (回車跳過): ",
	},
	"ja-jp": {
		"言語を選択してください(%d-%d): ",
		"Cookie と入力してください: ",
		"このCookieは正しくありません。再入力してください: ",
		"act_id（e先頭、長さ16）を入力してください: ",
		"%s (リターンズ・スキップ): ",
	},
	"ko-kr": {
		"언어 선택(%d-%d): ",
		"Cookie 를 입력하십시오: ",
		"이 쿠키가 올바르지 않습니다. 다시 입력하십시오: ",
		"act_id(e 시작, 길이 16)를 입력하십시오: ",
		"%s (캐리지 리턴 건너뛰기): ",
	},
	"en-us": {
		"Please select language (%d-%d): ",
		"Please enter Cookie: ",
		"The cookie is incorrect. Please re-enter: ",
		"Please enter act_id (start with e, length of 16): ",
		"%s (Enter to skip): ",
	},
}

var LangZoneMap = make(map[string]string)
var ZoneLangMap = make(map[string]string)

// GameNameMap 各语言游戏名称
var GameNameMap = map[string][]string{
	"zh-cn": {"原神", "崩坏：星穹铁道", "崩坏3rd", "未定事件簿"},
	"zh-tw": {"原神", "崩壞：星穹鐵道", "崩壞3rd", "未定事件簿"},
	"ja-jp": {"原神", "崩壊：スターレイル", "崩壊3rd", "未定事件簿"},
	"ko-kr": {"원신", "붕괴: 스타레일", "붕괴3rd", "미해결사건부"},
	"en-us": {"Genshin Impact", "Honkai: Star Rail", "Honkai Impact 3rd", "Tears of Themis"},
}

// SignMap 签到链接字典
var SignMap = map[string]string{
	"原神": "YS", "崩坏：星穹铁道": "XQTD", "崩坏3rd": "BH3", "未定事件簿": "WDSJB",
	"原神-tw": "YS", "崩壞：星穹鐵道": "XQTD", "崩壞3rd": "BH3", "未定事件簿-tw": "WDSJB",
	"原神-jp": "YS", "崩壊：スターレイル": "XQTD", "崩壊3rd": "BH3", "未定事件簿-jp": "WDSJB",
	"원신": "YS", "붕괴: 스타레일": "XQTD", "붕괴3rd": "BH3", "미해결사건부": "WDSJB",
	"Genshin Impact": "YS", "Honkai: Star Rail": "XQTD", "Honkai Impact 3rd": "BH3", "Tears of Themis": "WDSJB",
}

// ApiMap 签到接口列表
var ApiMap = map[string]string{
	"YS":    "https://sg-hk4e-api.hoyolab.com/event/sol/sign",                           // 原神
	"XQTD":  "https://sg-public-api.hoyolab.com/event/luna/os/sign",                     // 星穹铁道
	"BH3":   "https://sg-public-api.hoyolab.com/event/mani/sign",                        // 崩坏3rd
	"WDSJB": "https://sg-public-api.hoyolab.com/event/luna/os/sign",                     // 未定事件簿
	"TOKEN": "https://sg-public-api.hoyolab.com/account/ma-passport/token/verifyLToken", // 登录验证
}

var Lang = "English"       // 语言类型
var Zone = "en-us"         // 语言区域
var Interval = 8 * 60 * 60 // 执行间隔

func init() {
	for _, arr := range LangZoneList {
		LangZoneMap[arr[0]] = arr[1]
		ZoneLangMap[arr[1]] = arr[0]
	}
	shake := Interval / 10
	Interval = GetRandomInt(Interval-shake, Interval+shake)
}
