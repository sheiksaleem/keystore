package controllers
import (
	"github.com/beego/beego/v2/core/logs"
	"fmt"
	"encoding/json"
	"net/http"
)

type StoreKeyController struct {
	BaseController
}

type keyValue struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

func (c *StoreKeyController) GetKey() {
	log := logs.NewLogger()
	key := c.Ctx.Input.Query(":key")
	var err error
	var statusCode int
	fmt.Println("Here is the key", key)
	log.Debug(fmt.Sprintf("request key is :%s", key))

	defer func() {
		c.handleResponse(err, statusCode)
	}()

	value := c.GetKeyFromStore(key)
	log.Info(fmt.Sprintf("value is :%s", value))
	c.Data["JSON"] = map[string]string{key: value}
	c.ServeJSON()
}


func (c *StoreKeyController) SetKey() {
	log := logs.NewLogger()
	payload := c.Ctx.Input.RequestBody
	var err error
	var statusCode int

	var data keyValue
	fmt.Println("Here is the body", payload)

	defer func() {
		c.handleResponse(err, statusCode)
	}()

	err = json.Unmarshal([]byte(payload), &data)
	if err != nil {
		log.Error(fmt.Sprintf("Invalid Input payload: %s", err.Error()))
		statusCode = http.StatusBadRequest
		return
	}
	//value := c.GetKeyFromStore(key)
	c.SetKeyAtStore(data.Key, data.Value)
	c.Data["JSON"] = map[string]string{data.Key: data.Value}
	c.ServeJSON()
}


func (c *StoreKeyController) SearchKey() {
	log := logs.NewLogger()
	key := c.Ctx.Input.Query(":key")
	fmt.Println("Here is the key", key)
	log.Debug(fmt.Sprintf("request key is :%s", key))
	value := c.GetKeyFromStore(key)
	c.Data["JSON"] = map[string]string{key: value}
	c.ServeJSON()
}


func (c *StoreKeyController) handleResponse(err error, statusCode int) {
	if err != nil {
		c.Data["JSON"] = map[string]string{"error": err.Error()}
	}
	c.ServeJSON()
}
