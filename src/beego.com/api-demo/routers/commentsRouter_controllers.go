package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:DemoController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:DemoController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           `/index`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:DemoController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:DemoController"],
		beego.ControllerComments{
			Method:           "ManyToMany",
			Router:           `/postAndTag`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:DemoController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:DemoController"],
		beego.ControllerComments{
			Method:           "Set",
			Router:           `/session/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:DemoController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:DemoController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/session/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"],
		beego.ControllerComments{
			Method:           "Add",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:empno`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"],
		beego.ControllerComments{
			Method:           "Remove",
			Router:           `/:empno`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"],
		beego.ControllerComments{
			Method:           "Update",
			Router:           `/:empno`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"],
		beego.ControllerComments{
			Method:           "Query",
			Router:           `/queryTest`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"],
		beego.ControllerComments{
			Method:           "Transaction",
			Router:           `/transaction`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:EmpController"],
		beego.ControllerComments{
			Method:           "Validate",
			Router:           `/validate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego.com/api-demo/controllers:ToolController"] = append(beego.GlobalControllerRouter["beego.com/api-demo/controllers:ToolController"],
		beego.ControllerComments{
			Method:           "Lookup",
			Router:           `/profile/:type`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
