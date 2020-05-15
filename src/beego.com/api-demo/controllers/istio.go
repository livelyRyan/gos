package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"beego.com/api-demo/models"

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

// Operations about istio
type IstioController struct {
	beego.Controller
}

func (o *IstioController) Post() {

}

// @Title Get
// @Description find istio crd by skip
// @Param	page		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:page [get]
func (o *IstioController) Get() {
	skipStr := o.Ctx.Input.Param(":page")
	skip, err := strconv.ParseInt(skipStr, 10, 64)
	if err != nil {
		beego.Error("get skip data failed,", err, "\n")
		o.Data["json"] = err.Error()
		o.ServeJSON()
		return
	}
	limit, batch := int64(limit), int32(batchSize)
	skip = skip * int64(batch)
	results, err := models.Find(bson.D{}, &options.FindOptions{BatchSize: &batch, Limit: &limit, Skip: &skip})
	beego.Debug("result count ", len(results), "\n")
	o.Data["json"] = results
	o.ServeJSON()
}

// @Title GetAll
// @Description get all istio crd
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *IstioController) GetAll() {
	totalCount, err := models.EstimateDocumentCount()
	if err != nil {
		beego.Error("get data count failed,", err, "\n")
		o.Data["json"] = err.Error()
		o.ServeJSON()
		return
	}
	beego.Debug("total count ", totalCount, "\n")
	limit, skip, batch := int64(limit), int64(0), int32(batchSize)
	results, err := models.Find(bson.D{}, &options.FindOptions{BatchSize: &batch, Limit: &limit, Skip: &skip})
	beego.Debug("result count ", len(results), "\n")
	o.Data["json"] = results
	o.ServeJSON()
}

//DeleteOneVersion 删除一个版本的crd
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
// @router /addOne/:version [post]
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
	crd["metadata"] = interface{}(models.Metadata_{"nginx-gateway", models.Label{version}})
	_, err = models.InsertOne(crd, &options.InsertOneOptions{})
	if err != nil {
		o.Data["json"] = err.Error()
		o.ServeJSON()
		return
	}
	o.Data["json"] = "add one success"
	o.ServeJSON()
}
