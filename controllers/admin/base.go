package admin

import (
	"github.com/astaxie/beego"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
}

func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "admin"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)

}

// 显示模板,设置模板布局
func (this *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname = this.moduleName + "/" + tpl[0] + ".html"
	} else {
		tplname = this.moduleName + "/" + this.controllerName + "/" + this.actionName + ".html"
	}

	this.Data["version"] = beego.AppConfig.String("appversion")

	this.Layout = this.moduleName + "/layout.html"
	this.TplNames = tplname
}

func (this *BaseController) getTime() time.Time {
	return time.Now().UTC().Add(time.Duration(8 * float64(time.Hour)))
}
