package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	model "keystore/models"
	//"github.com/beego/beego/v2/core/logs"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BaseController struct {
	beego.Controller
}

func (c *BaseController) GetKeyFromStore(key string) string {
	return model.KeyValueStore.GetKey(key)
}

func (c *BaseController) SetKeyAtStore(key, value string) bool {
	return model.KeyValueStore.SetKey(key, value)
}

func (c *BaseController) SearchSuffixAtStore(key string) []string {
	return model.KeyValueStore.SearchSuffix(key)
}

func (c *BaseController) SearchPrefixAtStore(key string) []string {
	return model.KeyValueStore.SearchPrefix(key)
}
