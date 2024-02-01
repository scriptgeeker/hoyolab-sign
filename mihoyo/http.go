package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// SendPostRequest 发送 Post 请求
func SendPostRequest(url string, cookie string, jsonBody string) string {
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonBody))
	PrintError(err)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", cookie)
	Client := &http.Client{}
	resp, err := Client.Do(req)
	PrintError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		PrintError(err)
	}(resp.Body)
	all, err := io.ReadAll(resp.Body)
	jsonResp := string(all)
	return jsonResp
}

// GetResponse 获取响应信息
func GetResponse(jsonStr string) *ResponseVO {
	vo := ResponseVO{}
	err := json.Unmarshal([]byte(jsonStr), &vo)
	PrintError(err)
	return &vo
}
