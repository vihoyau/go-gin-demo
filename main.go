package main

import (
	"github.com/EDDYCJY/go-gin-demo/pkg/models"
	"github.com/EDDYCJY/go-gin-demo/pkg/setting"
)

// 初始化
func init() {
	// 初始化配置-> 需要创建setting文件-> 引入pkg文件夹 -> 创建setting 映射关系
	setting.Setup()
	//
	models.Setup()
}

func main() {
	//gin.SetMode()
}
