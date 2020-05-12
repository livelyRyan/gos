// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact bes@gmail.com
package routers

import (
	"beego.com/api-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/emp",
			beego.NSInclude(
				&controllers.EmpController{},
			),
		),
		beego.NSNamespace("/toolbox",
			beego.NSInclude(
				&controllers.ToolController{},
			),
		),
		beego.NSNamespace("/demo",
			beego.NSInclude(
				&controllers.DemoController{},
			),
		),
	)
	beego.AddNamespace(ns)

	// 添加过滤器
	/*var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		if !ok && ctx.Request.RequestURI != "/login" {
			ctx.Output.Status = http.StatusUnauthorized
			ctx.Output.Body([]byte("请进行登录"))
		}
	}
	beego.InsertFilter("/v1/demo/index", beego.BeforeRouter, FilterUser, true)*/
}
