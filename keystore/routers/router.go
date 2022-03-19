package routers

import (
	"keystore/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/keystore/get/?:key([\\w]+)", &controllers.StoreKeyController{}, "get:GetKey")
	beego.Router("/keystore/set", &controllers.StoreKeyController{}, "post:SetKey")
	beego.Router("/keystore/search", &controllers.StoreKeyController{}, "post:SearchKey")
}
