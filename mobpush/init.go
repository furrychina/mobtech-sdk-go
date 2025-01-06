package mobpush

import (
	"log"
	"os"
)

var (
	APPKey    = os.Getenv("MOB_PUSH_APP_KEY")
	APPSecret = os.Getenv("MOB_PUSH_APP_SECRET")
)

func init() {
	// 检查KEY和SECRET是否为空，为空则提示错误
	if APPKey == "" || APPSecret == "" {
		log.Fatal("MOB_PUSH_APP_KEY 或 MOB_PUSH_APP_SECRET 环境变量没有配置，请检查配置")
	}
}
