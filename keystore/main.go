package main

import (
	_ "keystore/routers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
)


func main() {
	beego.BConfig.Log.AccessLogs = true
	logs.SetLogger(logs.AdapterConsole)
	beego.Run("192.168.244.128")
} 


