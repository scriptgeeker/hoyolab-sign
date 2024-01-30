package main

// ResponseVO 通用响应结构
type ResponseVO struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Retcode int         `json:"retcode"`
}
