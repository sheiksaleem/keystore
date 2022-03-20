package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/spf13/cast"
	"net/http"
)

type StoreKeyController struct {
	BaseController
}

type keyValue struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func (c *StoreKeyController) GetKey() {
	key := c.Ctx.Input.Query(":key")
	var err error
	var statusCode int

	defer func() {
		c.handleResponse(err, statusCode)
	}()
	value := c.GetKeyFromStore(key)
	c.Data["json"] = &Response{Message: "Successfully retrived value", Data: map[string]string{key: value}}
}

func (c *StoreKeyController) SetKey() {
	log := logs.NewLogger()
	payload := c.Ctx.Input.RequestBody
	var err error
	var statusCode int
	var data keyValue

	defer func() {
		c.handleResponse(err, statusCode)
	}()

	err = json.Unmarshal([]byte(payload), &data)
	if err != nil {
		log.Error(fmt.Sprintf("Invalid Input payload: %s", err.Error()))
		statusCode = http.StatusBadRequest
		return
	}
	c.SetKeyAtStore(data.Key, cast.ToString(data.Value))
	c.Data["json"] = &Response{Message: "Successfully added key and value", Data: map[string]string{data.Key: cast.ToString(data.Value)}}
}

func (c *StoreKeyController) SearchKey() {
	queryParams := c.Ctx.Request.URL.Query()
	suffix := cast.ToString(queryParams.Get("suffix"))
	prefix := cast.ToString(queryParams.Get("prefix"))
	var result []string
	var err error
	var statusCode int

	defer func() {
		c.handleResponse(err, statusCode)
	}()
	if suffix != "" {
		result = c.SearchSuffixAtStore(suffix)
		fmt.Println(result)
	} else if prefix != "" {
		result = c.SearchPrefixAtStore(prefix)

	}
	c.Data["json"] = &Response{Message: "Successfully added key and value", Data: result}
}

func (c *StoreKeyController) handleResponse(err error, statusCode int) {
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	}
	c.ServeJSON()
}
