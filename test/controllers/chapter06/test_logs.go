package chapter06

import (
	"github.com/astaxie/beego/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type LogsController struct {
	beego.Controller
}

func (l *LogsController) Get() {

	//从上往下依次降低
	logs.Emergency("这是Emergency日志")
	logs.Alert("这是Alert日志")
	logs.Critical("这是Critical日志")
	logs.Error("这是Error日志")
	logs.Warn("这是Warn日志")
	logs.Notice("这是Notice日志")
	logs.Info("这是Info日志")
	logs.Debug("这是debug日志")

	l.TplName = "chapter06/test_logs.html"
}
