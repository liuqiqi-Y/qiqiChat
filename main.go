package main

import (
	"qiqiChat/conf"
	"qiqiChat/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run("172.16.4.178:3000")
}
