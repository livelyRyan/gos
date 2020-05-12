package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

type ToolController struct {
	beego.Controller
}

// @router /profile/:type [GET]
func (c *ToolController) Lookup() {
	logs.Info("enter Lookup function")
	toolbox.ProcessInput("lookup "+c.GetString(":type"), c.Ctx.ResponseWriter)
}
