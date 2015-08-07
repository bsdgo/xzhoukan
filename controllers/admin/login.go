package admin

import ()

type LoginController struct {
	BaseController
}

func (c *LoginController) Login() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "admin/login/login.html"
}
