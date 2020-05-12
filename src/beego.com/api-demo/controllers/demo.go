package controllers

import (
	"beego.com/api-demo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type DemoController struct {
	beego.Controller
}

// @router /postAndTag [GET]
func (c *DemoController) ManyToMany() {
	var posts []*models.Post
	_, err := orm.NewOrm().QueryTable("post").Filter("Tags__Tag__Name", "1").All(&posts)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = posts
	}
	c.ServeJSON()
}

// @router /session/:key [post]
func (c *DemoController) Set() {
	c.SetSession(c.GetString(":key", "defaultKey"), c.GetString("value", "defaultValue"))
	c.Data["json"] = "success"
	c.ServeJSON()
}

// @router /session/:key [get]
func (c *DemoController) Get() {
	val := c.GetSession(c.GetString(":key", "defaultKey"))
	c.Data["json"] = val
	c.ServeJSON()
}

// @Success 200
// @router /index [get]
func (c *DemoController) Index() {
	logs.Info("enter Index function")
	c.Data["json"] = "this is the index page"
	c.ServeJSON()
}
