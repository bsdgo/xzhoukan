package routers

import (
	"github.com/astaxie/beego"
	"xzhoukan.com/controllers"
	"xzhoukan.com/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/admin/login", &admin.LoginController{}, "*:Login")
}
