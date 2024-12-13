package global

import (
	"time"
)

const DEBUG = false //控制程序是否处于调试模式

const API = "localhost:5002" //指定服务监听的IP地址和端口

const HookURL = ""    //存储Webhook的URL地址
const LogDirPath = "" //指定日志文件的存储路径

const ProSalt = "w2t3f4d1304c381a" //pro盐值
const JWTSalt = "That1is9a7secret" //JSON盐值

var CurrentVersion = Version("1.2.2")              //当前程序版本号
var CurrentVersionCode = VersionCode("2024_07_14") //当前程序版本代码

const UpdateSourceURL = ""         // 自动更新地址
const MaxRequestNumPerSecond = 100 //每秒限制的uida请求

/** 时间间隔常量 **/
const ProInterval = time.Second * 2      //进程或任务的执行间隔
const SwitchInterval = time.Second * 3   //控制某个状态切换或切换操作的间隔时间
const VersionInterval = time.Second / 10 //控制版本检查或更新的频率
const OnlineInterval = time.Second * 550 //程序检查在线状态或执行在线相关任务的间隔时间
const OrderInterval = time.Second * 3    //控制订单处理或订单相关任务的执行间隔
const OrderTimeoutExpress = "5m"         //订单的超时时间为5分钟
const OrderDuration = time.Minute * 5    // 订单的有效持续时间或订单处理的总时长

const ChargeModeRate = 7                  //指定计费模式的比率
const ChargeFeeSecond = time.Minute * 171 //指定计费周期

// Postgres 数据库
const PostgresIP = ""
const PostgresUsr = ""
const PostgresPwd = ""

const PostgresDB = ""

func init() {
	if DEBUG {

	}
}
