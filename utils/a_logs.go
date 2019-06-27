package utils

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var Logger *logs.BeeLogger

// 保证此文件最早加载
// 初始化日志
func init() {
	runmode := strings.TrimSpace(strings.ToLower(beego.AppConfig.DefaultString("runmode", "dev")))
	if runmode == "dev" {
		Logger = logs.NewLogger(1)
		Logger.SetLogger(logs.AdapterConsole)
		// 异步
		Logger.Async()
	} else {
		Logger = logs.NewLogger(1000)
		level := beego.AppConfig.String("logs.level")
		Logger.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/`+runmode+`.log",
			"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
			"level":`+level+`,
			"daily":true,
			"maxdays":10}`)
		// 异步
		Logger.Async(1e3)
	}
	Logger.Info("Init: Logger: %s", Logger)
}
