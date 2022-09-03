package util

import "github.com/EDDYCJY/go-gin-demo/pkg/setting"

func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
