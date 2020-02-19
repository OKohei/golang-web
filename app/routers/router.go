package routers

import (
	"github.com/app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/students", &controllers.FirstController{}, "get:GetStudents")
    beego.Router("/dashboard", &controllers.FirstController{}, "get:Dashboard")
}
