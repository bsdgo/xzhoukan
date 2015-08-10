package routers

import (
	"github.com/astaxie/beego"
	"xzhoukan.com/controllers"
	"xzhoukan.com/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/admin/login", &admin.LoginController{}, "get:Login")
	beego.Router("/admin/login", &admin.LoginController{}, "post:DoLogin")
	beego.Router("/admin", &admin.MainController{}, "get:Index")

	beego.Router("/admin/issue/add", &admin.IssueController{}, "get:Add")
	beego.Router("/admin/issue/add", &admin.IssueController{}, "post:DoAdd")
}
