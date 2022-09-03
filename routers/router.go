package routers

import (
	"github.com/EDDYCJY/go-gin-demo/pkg/export"
	"github.com/EDDYCJY/go-gin-demo/pkg/qrcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 使用框架的中间件
	// 日志 -> 包含状态码、响应时长等信息 -> 此处可以重写，根据自己的需要返回 -> 例如分布式链路的信息
	r.Use(gin.Logger())
	//  代码中出现panic恐慌 -> 捕获异常，并且返回500
	r.Use(gin.Recovery())

	// 获取静态文件, 加载所有的静态链路, 方便返回
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	return r

}
