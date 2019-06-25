package routers

import (
	"player/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/query", &controllers.QueryController{})
	beego.Router("/play", &controllers.PlayController{})

}
