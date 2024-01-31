# hoyolab-sign

hoyolab 社区自动签到，领取游戏奖励
Hoyolab Community automatic sign-in to receive game rewards

### 支持游戏
- 原神          Genshin Impact
- 崩坏：星穹铁道 Honkai: Star Rail
- 崩坏3rd       Honkai Impact 3rd
- 未定事件簿    Tears of Themis

### 支持语言
- 简体中文  zh-cn
- 繁体中文  zh-tw
- 日本語    ja-jp
- 한국어    ko-kr
- English   en-us

### 项目结构

- conf.go   全局配置
- file.go   文件操作模块
- go.mod    项目模块
- http.go   网络请求模块
- main.go   程序入口
- serv.go   函数方法模块
- tick.go   定时器模块
- tool.go   封装工具
- view.go   结构实体

### 使用指导

0. 运行程序，控制台执行 `go run mihoyo`，也可直接使用 release 中已编译好的可执行文件
1. 选择语言，该选项会影响签到结果，请根据游戏服务所在区域选择对应语言
2. 输入 Cookie，建议使用浏览器插件 Cookie-Editor 导出 Head String 填入
3. 输入 act_id，需要从对应签到页面的地址栏获取，类似 
4. 自动签到，程序会每隔 8 小时自动完成签到
