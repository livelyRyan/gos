package controllers

import (
	"beego.com/api-demo/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

// Operations about Emps
type EmpController struct {
	beego.Controller
}

var ch cache.Cache

func init() {
	ch = cache.NewMemoryCache()
}

// @Title Get
// @Description find object by empno
// @Param	empno		path 	string	true		"the empno you want to get"
// @Success 200 {object} models.Emp
// @Failure 403 :objectId is empty
// @router /:empno [get]
func (c *EmpController) Get() {
	empno := c.Ctx.Input.Param(":empno")
	if empno != "" {
		// 从缓存中获取
		emp := ch.Get(empno)
		if emp != nil {
			c.Data["json"] = emp
			c.ServeJSON()
			return
		}

		id, err := strconv.ParseInt(empno, 10, 32)
		if err != nil {
			c.Data["json"] = err.Error()
			c.ServeJSON()
			return
		}
		emp = &models.Emp{Id: int32(id)}
		err = orm.NewOrm().Read(emp)

		if err == nil {
			// 存入缓存
			err = ch.Put(empno, emp, time.Hour)
			if err != nil {
				logs.Error("put cache value error: %v", err)
			}
		}
		setReturnJson(c, emp, err)
	}
}

// @Success 200 {object} models.Emp
// @router / [get]
func (c *EmpController) GetAll() {
	var emps []*models.Emp
	_, err := orm.NewOrm().QueryTable(&models.Emp{}).All(&emps)
	setReturnJson(c, emps, err)
}

// @router / [post]
func (c *EmpController) Add() {
	emp := new(models.Emp)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, emp)
	if err != nil {
		c.Data["json"] = err.Error()
		return
	}

	_, err = orm.NewOrm().Insert(emp)
	setReturnJson(c, emp, err)
}

// @Success 200
// @router /:empno [delete]
func (c *EmpController) Remove() {
	empno, err := c.GetInt32(":empno")
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	emp := models.Emp{Id: empno}
	_, err = orm.NewOrm().Delete(&emp)
	setReturnJson(c, empno, err)
}

// @Success 200
// @router /:empno [put]
func (c *EmpController) Update() {
	emp := new(models.Emp)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, emp)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	r, err := orm.NewOrm().Update(emp)
	setReturnJson(c, r, err)
}

// @Failure 400 :emp params invalid
// @router /validate [post]
func (c *EmpController) Validate() {
	var emp models.Emp
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &emp)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		errs := models.Validate(&emp)
		c.Ctx.Output.Status = http.StatusBadRequest
		if errs != nil {
			c.Data["json"] = errs
		} else {
			c.Data["json"] = "correct format"
		}
	}
	c.ServeJSON()
}

// @router /transaction [GET]
func (c *EmpController) Transaction() {
	var errors []error
	ormer := orm.NewOrm()

	defer func() {
		var err error
		if len(errors) > 0 {
			err = ormer.Rollback()
		} else {
			err = ormer.Commit()
		}
		if err != nil {
			errors = append(errors, err)
		}
		setReturnJson(c, "success", errors)
	}()

	err := ormer.Begin()
	if err != nil {
		errors = append(errors, err)
		return
	}
	// 事务处理过程
	// 此过程中的所有使用 ormer 对象的查询也会算在事务内
	insertToDeptSQL := "insert into dept values (?, ?, ?)"
	_, err = ormer.Raw(insertToDeptSQL, 77, "TEST", "BEIJING").Exec()
	if err != nil {
		errors = append(errors, err)
		return
	}
	_, err = ormer.Raw("insert into emp(empno, ename, deptno) values(7777, 'clearlove', 77)").Exec()
	if err != nil {
		errors = append(errors, err)
		return
	}
}

// @router /queryTest [get]
func (c *EmpController) Query() {
	typ, err := c.GetInt64("type", 0)
	if err != nil {
		c.Data["json"] = err.Error()
		return
	}
	var data interface{}
	switch typ {
	case 1:
		data, err = models.QueryAllBySQL()
	case 2:
		data, err = models.QueryValues()
	case 3:
		data, err = models.QueryBuilder()
	case 4:
		data, err = models.QuerySeter()
	case 6:
		data, err = models.OneToMany()
	case 7:
		data, err = models.ManyToOne()
	default:
		data = models.Emp{Id: 666}
	}
	setReturnJson(c, data, err)
}

func setReturnJson(c *EmpController, data interface{}, err interface{}) {
	if err != nil && (reflect.TypeOf(err).Kind() != reflect.Slice || reflect.ValueOf(err).Len() > 0) {
		c.Data["json"] = err
	} else {
		c.Data["json"] = data
	}
	c.ServeJSON()
}
