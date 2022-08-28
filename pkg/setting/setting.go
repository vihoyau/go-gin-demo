package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
type DataBase struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var ServerSetting = &Server{}
var AppSetting = &App{}
var DataBaseSetting = &DataBase{}
var RedisSetting = &Redis{}

// 声明全局变量，引入ini库
var cfg *ini.File

func Setup() {
	var err error
	// 加载配置文件-> 编写配置文件（应用环境变量、文件大小、服务端口及读写超时限制、数据库及表、redis配置）
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		// 采取Fatal的形式，规范报错信息。
		log.Fatalf("Setting.Setup, fail to parse 'conf/app.ini: %v", err)
	}
	// 映射结构体的关系-> 好像java的getter、setter基础类
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DataBaseSetting)
	mapTo("redis", RedisSetting)

	// 改造数据单位
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	// 获取配合信息及映射关系，判断映射过程是否报错。
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo to %s, err: %v", section, err)
	}
}
