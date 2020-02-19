package routers

import (
	"github.com/app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/students", &controllers.FirstController{}, "get:GetStudents")
    beego.Router("/dashboard", &controllers.FirstController{}, "get:Dashboard")

	beego.Router("/home", &controllers.SessionController{}, "get:Home")
    beego.Router("/login", &controllers.SessionController{}, "get:Login")
    beego.Router("/logout", &controllers.SessionController{}, "get:Logout")
}
