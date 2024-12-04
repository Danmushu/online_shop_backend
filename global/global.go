package global

import (
	"time"
)

const DEBUG = false

const API = "0.0.0.0:5002"

//const API = "0.0.0.0:80"

const HookURL = "https://albion.scoreflow.cn/api/hook"
const LogDirPath = "./log"

// const ProSalt = "w1t4f3z1304c381a" // 不要修改
const ProSalt = "w2t3f4d1304c381a"
const JWTSalt = "That1is9a7secret"

var CurrentVersion = Version("1.2.2")
var CurrentVersionCode = VersionCode("2024_07_14")

const UpdateSourceURL = "" // 自动更新地址
const MaxRequestNumPerSecond = 100

const ProInterval = time.Second * 2
const SwitchInterval = time.Second * 3
const VersionInterval = time.Second / 10
const OnlineInterval = time.Second * 550
const OrderInterval = time.Second * 3
const OrderTimeoutExpress = "5m"
const OrderDuration = time.Minute * 5

const ChargeModeRate = 7
const ChargeFeeSecond = time.Minute * 171

// PostgresAPI 数据库
const PostgresAPI = "47.94.225.169"
const PostgresUsr = "ydx"
const PostgresPwd = "yangdongxu"

// const PostgresDB = "base"
// const PostgresDB = "hcnbg6288@outlook.com"
const PostgresDB = "2292871544@qq.com"

func init() {
	if DEBUG {

	}
}
