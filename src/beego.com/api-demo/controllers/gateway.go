package controllers

import (
	"beego.com/api-demo/models"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

//  GatewayController operations for Gateway
type GatewayController struct {
	beego.Controller
}

// URLMapping ...
func (c *GatewayController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Gateway
// @Param	body		body 	models.Gateway	true		"body for Gateway content"
// @Success 201 {int} models.Gateway
// @Failure 403 body is empty
// @router / [post]
func (c *GatewayController) Post() {
	var v models.Gateway
	setTestValue(&v)
	//json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.AddGateway(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func setTestValue(gateway *models.Gateway) {
	gateway.Metadata.Namespace = "test"
	gateway.Metadata.Name = strconv.FormatInt(time.Now().Unix(), 10)
	server := new(models.Server)
	server.Tls = &models.ServerTLSSettings{
		ServerCertificate: "testCert",
		PrivateKey:        "testKey",
	}
	server.Port = &models.Port{
		Number:   80,
		Protocol: "http",
		Name:     "http",
	}

	gateway.Spec.Servers = append(gateway.Spec.Servers, server)
}

// GetOne ...
// @Title Get One
// @Description get Gateway by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Gateway
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GatewayController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetGatewayById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Gateway
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Gateway
// @Failure 403
// @router / [get]
func (c *GatewayController) GetAll() {
	/*var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64*/

	var fields []string
	sortby := []string{"name"}
	var order []string
	var query = make(map[string]string)
	query["namespace"] = "test"
	limit := int64(2)
	offset := int64(0)

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllGateway(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Gateway
// @Param	namespace		path 	string	true		"The gateway resource namespace you want to delete"
// @Param	name		path 	string	true		"The gateway resource name you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 namespace or name is empty
// @router /:namespace/:name [delete]
func (c *GatewayController) Delete() {
	defer c.ServeJSON()
	ns := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":name")
	if ns == "" || name == "" {
		c.Ctx.ResponseWriter.Status = http.StatusForbidden
		c.Data["json"] = "parameter 'namespace' or 'name' is empty"
		return
	}

	if err := models.DeleteGateway(ns, name); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
}
