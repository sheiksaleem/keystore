package controllers

import (
	beego "github.com/beego/beego/v2/server/web"	
	model "keystore/models"
	"github.com/beego/beego/v2/core/logs"
)

type BaseController struct {
	beego.Controller
}


func (c *BaseController) GetKeyFromStore(key string) string {
	log := logs.NewLogger()
	log.Debug("Getting key value from store")
	return model.KeyValueStore.GetKey(key)
}



func (c *BaseController) SetKeyAtStore(key, value string) bool {
	log := logs.NewLogger()
	log.Debug("Setting key and value at store")
	return model.KeyValueStore.SetKey(key, value)
}