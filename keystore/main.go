package main

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "keystore/routers"
)

func main() {
	beego.BConfig.Log.AccessLogs = true
	logs.SetLogger(logs.AdapterConsole)

	beego.Run(beego.BConfig.Listen.HTTPAddr)
}
