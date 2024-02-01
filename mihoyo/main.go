package main

import (
	"fmt"
)

func main() {
	fmt.Println("====== HoYoLab Auto Sign ======")
	Lang = LoadLang()     // 加载语言选项
	Cookie = LoadCookie() // 校验 Cookie
	UpdateGamesMap()      // 更新游戏字典
	TimedSign()           // 开启定时任务
}
