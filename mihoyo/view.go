package main

// ResponseVO 通用响应结构
type ResponseVO struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Retcode int         `json:"retcode"`
}

// SignBodyVO 签到请求体结构
type SignBodyVO struct {
	ActId string `json:"act_id"`
	Lang  string `json:"lang"`
}

// LangVO 语言信息结构
type LangVO struct {
	Name string
	Lang string
}

// GameVO 游戏信息结构
type GameVO struct {
	Name    string
	ActId   string
	SignUrl string
}
