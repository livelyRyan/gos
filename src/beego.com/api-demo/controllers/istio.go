package controllers

import (
	"beego.com/api-demo/models"
	"errors"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	batchSize = 20
	limit     = 20
)

func init() {
	//models.InitDatabaseData()
}

//IstioController istio
// Operations about istio
type IstioController struct {
	beego.Controller
}

//Post post function
func (o *IstioController) Post() {

}

//Get 通过指定参数数据获取一条数据
// @Title Get
// @Description find istio crd by skip
// @Param	version		path 	string	true		"The version of crd to query"
// @Success 200 {object} crd
// @Failure 403 :version is empty
// @router /:version [get]
func (o *IstioController) Get() {
	defer o.ServeJSON()
	versionStr := o.Ctx.Input.Param(":version")
	version, err := strconv.ParseInt(versionStr, 10, 64)
	if err != nil {
		beego.Error("get skip data failed,", err, "\n")
		o.Data["json"] = err.Error()
		return
	}
	versionStr = fmt.Sprintf("v%d", version)
	filter := bson.D{
		{
			"metadata", bson.D{
				{"name", "nginx-gateway"},
				{"labels", bson.D{
					{"version", versionStr},
				}},
			},
		},
	}
	result, err := models.FindOne(filter)
	if err != nil {
		beego.Error("get data from mongo failed,", err, "\n")
		o.Data["json"] = err.Error
		return
	}
	o.Data["json"] = result
}

//GetAll 获取数据
// @Title GetAll
// @Description get all istio crd
// @Param	page	query	string	false	"which page to show. Must be an integer"
// @Success 200 {object} crd
// @Failure 403 :objectId is empty
// @router / [get]
func (o *IstioController) GetAll() {
	defer o.ServeJSON()
	page := o.Ctx.Input.Query("page")
	skip := int64(0)
	if len(page) > 0 {
		tempSkip, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error("get page data failed,", err, "\n")
			o.Data["json"] = err.Error()
			return
		}
		skip = tempSkip * int64(batchSize)
	}
	limit, batch := int64(limit), int32(batchSize)
	results, err := models.Find(bson.D{}, &options.FindOptions{BatchSize: &batch, Limit: &limit, Skip: &skip})
	if err != nil {
		beego.Error("get data from mongo failed,", err, "\n")
		o.Data["json"] = err.Error
		return
	}
	beego.Debug("result count ", len(results), "\n")
	o.Data["json"] = results
}

//DeleteOneByVersion 删除一个版本的crd
func (o *IstioController) DeleteOneByVersion() {
	version := o.Ctx.Input.Param(":version")
	if len(version) == 0 {
		o.Data["json"] = errors.New("version is empty")
		o.ServeJSON()
		return
	}
	version = fmt.Sprintf("v%s", version)
	filter := bson.D{
		{
			"metadata", bson.D{
				{"name", "nginx-gateway"},
				{"labels", bson.D{
					{"version", version},
				}},
			},
		},
	}
	result, err := models.DeleteOne(filter, nil)
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
		return
	}
	if result.DeletedCount == 0 {
		o.Data["json"] = "no match crd to delete"
		o.ServeJSON()
		return
	}
	o.Data["json"] = "delete success"
	o.ServeJSON()
}

//AddOneByVersion 增加一个版本的crd
func (o *IstioController) AddOneByVersion() {
	version := o.Ctx.Input.Param(":version")
	if len(version) == 0 {
		o.Data["json"] = errors.New("version is empty")
		o.ServeJSON()
		return
	}
	version = fmt.Sprintf("v%s", version)
	crd, err := models.GetYaml()
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
		return
	}
	crd["metadata"] = bson.D{{"name", "nginx-gateway"}, {"labels", bson.M{"version": version}}}
	_, err = models.InsertOne(crd, &options.InsertOneOptions{})
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
		return
	}
	o.Data["json"] = "add one success"
	o.ServeJSON()
}
