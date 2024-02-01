# hoyolab-sign

hoyolab 社区自动签到，领取每日奖励

Hoyolab Community automatic sign-in

### 游戏列表

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

### 生成文件

- lang.cnf      语言设置，清空可重新设置语言选项
- cookie.txt    登录状态，用于保存用户的 Cookie
- act-id.cnf    活动ID表，如果活动过期可手动配置
- sign.log      签到日志，用于记录签到接口调用情况

### 使用指导

1. 运行程序，直接运行 release 中已编译好的可执行文件
2. 选择语言，请根据游戏服务所在区域选择对应语言
3. 输入 Cookie，使用 Cookie-Editor 导出 Header String 填入
4. 自动签到，程序会每隔 8 小时为所有游戏签到一次

