package main

// COOKIE_PATH cookie 配置文件
const COOKIE_PATH = "./temp/cookie.txt"

// LANG_PATH 语言配置文件
const LANG_PATH = "./temp/lang.cnf"

// ACT_IDS_PATH 角色id列表
const ACT_IDS_PATH = "./temp/act-id.cnf"

// SIGN_LOG_PATH 签到日志记录
const SIGN_LOG_PATH = "./temp/sign.log"

// LangList 语言信息列表
var LangList = []LangVO{
	{"简体中文", "zh-cn"},
	{"繁體中文", "zh-tw"},
	{"日本語", "ja-jp"},
	{"한국어", "ko-kr"},
	{"English", "en-us"},
}

// TipMap 多语言提示词
var TipMap = map[string][]string{
	"zh-cn": {
		"请选择语言 (%d-%d): ",
		"请输入 Cookie: ",
		"该 Cookie 不正确, 请重新输入: ",
		"每日签到完成",
	},
	"zh-tw": {
		"請選擇語言 (%d-%d): ",
		"請輸入 Cookie: ",
		"該 Cookie 不正確, 請重新輸入: ",
		"每日簽到完成",
	},
	"ja-jp": {
		"言語を選択してください(%d-%d): ",
		"Cookie と入力してください: ",
		"このCookieは正しくありません。再入力してください: ",
		"毎日のサインアップ完了",
	},
	"ko-kr": {
		"언어 선택(%d-%d): ",
		"Cookie 를 입력하십시오: ",
		"이 쿠키가 올바르지 않습니다. 다시 입력하십시오: ",
		"매일 출석 완료",
	},
	"en-us": {
		"Please select language (%d-%d): ",
		"Please enter Cookie: ",
		"The cookie is incorrect. Please re-enter: ",
		"Daily sign-in completed",
	},
}

// ApiMap 签到接口列表
var ApiMap = map[string]string{
	"GI":    "https://sg-hk4e-api.hoyolab.com/event/sol/sign",                           // 原神
	"HSR":   "https://sg-public-api.hoyolab.com/event/luna/os/sign",                     // 星穹铁道
	"HI3":   "https://sg-public-api.hoyolab.com/event/mani/sign",                        // 崩坏3rd
	"TOT":   "https://sg-public-api.hoyolab.com/event/luna/os/sign",                     // 未定事件簿
	"TOKEN": "https://sg-public-api.hoyolab.com/account/ma-passport/token/verifyLToken", // 登录验证
}

// GamesMap 游戏区服信息
var GamesMap = map[string][]GameVO{
	"zh-cn": {
		{"原神", "e202102251931481", ApiMap["GI"]},
		{"崩坏：星穹铁道", "e202303301540311", ApiMap["HSR"]},
		{"崩坏3rd", "e202110291205111", ApiMap["HI3"]},
		{"未定事件簿", "e202202281857121", ApiMap["TOT"]},
	},
	"zh-tw": {
		{"原神", "e202102251931481", ApiMap["GI"]},
		{"崩壞：星穹鐵道", "e202303301540311", ApiMap["HSR"]},
		{"崩壞3rd", "e202110291205111", ApiMap["HI3"]},
		{"未定事件簿", "e202308141137581", ApiMap["TOT"]},
	},
	"ja-jp": {
		{"原神", "e202102251931481", ApiMap["GI"]},
		{"崩壊：スターレイル", "e202303301540311", ApiMap["HSR"]},
		{"崩壊3rd", "e202110291205111", ApiMap["HI3"]},
		{"未定事件簿", "e202202281857121", ApiMap["TOT"]},
	},
	"ko-kr": {
		{"원신", "e202102251931481", ApiMap["GI"]},
		{"붕괴: 스타레일", "e202303301540311", ApiMap["HSR"]},
		{"붕괴3rd", "e202110291205111", ApiMap["HI3"]},
		{"미해결사건부", "e202202281857121", ApiMap["TOT"]},
	},
	"en-us": {
		{"Genshin Impact", "e202102251931481", ApiMap["GI"]},
		{"Honkai: Star Rail", "e202303301540311", ApiMap["HSR"]},
		{"Honkai Impact 3rd", "e202110291205111", ApiMap["HI3"]},
		{"Tears of Themis", "e202202281857121", ApiMap["TOT"]},
	},
}

var Lang = "en-us" // 语言区域
var Cookie string
var Interval = 8 * 60 * 60 // 执行间隔
