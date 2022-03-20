package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"keystore/controllers"
)

func init() {
	beego.Router("/keystore/get/?:key(.+)", &controllers.StoreKeyController{}, "get:GetKey")
	beego.Router("/keystore/set", &controllers.StoreKeyController{}, "post:SetKey")
	beego.Router("/keystore/search", &controllers.StoreKeyController{}, "post:SearchKey")
}
