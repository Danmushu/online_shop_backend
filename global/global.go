package global

import "time"

const DEBUG = false //控制程序是否处于调试模式

const API = "localhost:5002" //指定服务监听的IP地址和端口

const HookURL = ""         //存储Webhook的URL地址
const LogDirPath = "./log" //指定日志文件的存储路径

const JWTDuration = time.Hour * 24
const JWTSecret = ""

func init() {
	if DEBUG {

	}
}
