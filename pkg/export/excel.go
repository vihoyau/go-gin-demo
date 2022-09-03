package export

import "github.com/EDDYCJY/go-gin-demo/pkg/setting"

const EXT = ".xlsx"

func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}

// GetExcelPath 获取Excel相对目录地址
func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

// GetExcelFullPath 获取Excel绝对目录地址
func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}
