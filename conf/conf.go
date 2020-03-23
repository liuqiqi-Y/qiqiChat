package conf

import (
	"os"
	"qiqiChat/model"
	"qiqiChat/util"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load("env")

	// 设置日志级别
	util.SetLogLevel(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Err.Fatalln("翻译文件加载失败", err)
	}

	// 连接数据库
	//fmt.Println("===>", os.Getenv("MYSQL_DSN"))
	model.InitDB(os.Getenv("MYSQL_DSN"))
	//cache.Redis()
}
