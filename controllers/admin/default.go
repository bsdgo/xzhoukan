package admin

import (
//"fmt"
)

type MainController struct {
	BaseController
}

func (this *MainController) Index() {
	//this.Ctx.WriteString("hello world")

	//this.TplNames = "admin/login/login.html"
	this.display()
}
