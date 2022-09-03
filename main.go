package main

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-demo/pkg/gredis"
	"github.com/EDDYCJY/go-gin-demo/pkg/logging"
	"github.com/EDDYCJY/go-gin-demo/pkg/models"
	"github.com/EDDYCJY/go-gin-demo/pkg/setting"
	"github.com/EDDYCJY/go-gin-demo/pkg/util"
	"github.com/EDDYCJY/go-gin-demo/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 初始化
func init() {
	// 初始化配置-> 需要创建setting文件-> 引入pkg文件夹 -> 创建setting 映射关系
	setting.Setup()
	// 初始化数据库-> 创建models数据库表映射->
	models.Setup()
	// 初始化日志打印 -> 捕获运行时的报错信息 —> 做好报错格式处理 -> 按等级划分
	logging.Setup()
	// redis 数据库
	gredis.Setup()
	// 初始化JWT鉴权密钥
	util.Setup()
}

func main() {
	// 运行时的环境 -> 选择debug模式
	gin.SetMode(setting.ServerSetting.RunMode)
	// 初始化路由 -> 所有的请求链路都将写入这里
	routersInit := routers.InitRouter()
	// 设置下行超时限制
	readTimeout := setting.ServerSetting.ReadTimeout
	// 设置上行超时限制
	writeTimeout := setting.ServerSetting.WriteTimeout
	// 对端口打印并赋值
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	// 设置最大请求头数据量，左移20个字符。相当于100000000000000...
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
