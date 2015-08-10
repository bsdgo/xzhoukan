package admin

import (
	"fmt"
	"strings"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "admin/login/login.html"
}

func (this *LoginController) DoLogin() {
	username := strings.TrimSpace(this.GetString("username"))
	password := strings.TrimSpace(this.GetString("password"))

	fmt.Println(username + password)

	if username == "bozz" {
		this.Redirect("/admin", 302)
	}

	this.EnableRender = false
}
