package main

import (
	"github.com/EDDYCJY/go-gin-demo/pkg/logging"
	"github.com/EDDYCJY/go-gin-demo/pkg/models"
	"github.com/EDDYCJY/go-gin-demo/pkg/setting"
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
	//gredis.Setup()
}

func main() {
	//gin.SetMode()
}
