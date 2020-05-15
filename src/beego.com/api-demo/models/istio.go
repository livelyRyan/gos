package models

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

const (
	dbName         = "bes"
	collectionName = "istio"
)

func getCollection() (*mongo.Collection, error) {
	client, err := getMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(dbName).Collection(collectionName)
	return collection, nil
}

//Find 查找多条符合条件的数据
func Find(filter interface{}, opts ...*options.FindOptions) ([]bson.M, error) {
	collection, err := getCollection()
	if err != nil {
		beego.Error("get mongo client failed,", err, "\n")
		return nil, err
	}
	cursor, err := collection.Find(context.TODO(), filter, opts...)
	if err != nil {
		beego.Error("find mongo data failed,", err, "\n")
		return nil, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		beego.Error("parse result failed,", err, "\n")
		return nil, err
	}
	return results, nil
}

//FindOne 查找符合条件的一条数据
func FindOne(filter interface{}, opts ...*options.FindOneOptions) (bson.M, error) {
	collection, err := getCollection()
	if err != nil {
		beego.Error("get mongo client failed,", err, "\n")
		return nil, err
	}
	var result bson.M
	err = collection.FindOne(context.TODO(), filter, opts...).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

//InsertOne 新增一条数据
func InsertOne(document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	collection, err := getCollection()
	if err != nil {
		beego.Error("get mongo client failed,", err, "\n")
		return nil, err
	}
	result, err := collection.InsertOne(context.Background(), document, opts...)
	if err != nil {
		beego.Error("insert to mongo failed,", err, "\n")
		return nil, err
	}
	beego.Debug("insert result is ", result, "\n")
	return result, nil
}

//DeleteOne 删除一条数据
func DeleteOne(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	collection, err := getCollection()
	if err != nil {
		beego.Error("get mongo client failed,", err, "\n")
		return nil, err
	}
	result, err := collection.DeleteOne(context.TODO(), filter, opts...)
	if err != nil {
		beego.Error("delete mongo data failed,", err, "\n")
		return nil, err
	}
	return result, nil
}

//FindOneAndUpdate 找到一条数据并且更新
func FindOneAndUpdate(filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) (bson.M, error) {
	collection, err := getCollection()
	if err != nil {
		beego.Error("get mongo client failed,", err, "\n")
		return nil, err
	}
	var result bson.M
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, opts...).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		beego.Error("find and update failed,", err, "\n")
		return nil, err
	}
	return result, nil
}

//EstimateDocumentCount 得到估算的数量
func EstimateDocumentCount(opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	collection, err := getCollection()
	if err != nil {
		beego.Error("get mongo client failed,", err, "\n")
		return 0, err
	}
	return collection.EstimatedDocumentCount(context.Background(), opts...)
}

//GetYaml 从文件中获取yaml数据
func GetYaml() (map[string]interface{}, error) {
	buffer, err := ioutil.ReadFile("./crd/gateway.yaml")
	if err != nil {
		beego.Error("read file failed,", err, "\n")
		return nil, err
	}
	t := map[string]interface{}{}
	err = yaml.Unmarshal(buffer, &t)
	if err != nil {
		beego.Error("unmarshal failed,", err, "\n")
		return nil, err
	}
	return t, nil
}

//InitDatabaseData 初始化数据库数据，当前验证阶段造10000条数据
func InitDatabaseData() error {
	t, err := GetYaml()
	if err != nil {
		return err
	}
	collection, err := getCollection()
	if err != nil {
		beego.Error("get mongo client failed,", err, "\n")
		return err
	}

	for i := 0; i < 10000; i++ {
		t["metadata"] = bson.D{{"name", "nginx-gateway"}, {"labels", bson.M{"version": fmt.Sprintf("v%d", i)}}}
		_, err = collection.InsertOne(context.Background(), t)
		if err != nil {
			beego.Error("insert data to mongo failed,", err, "\n")
			return err
		}
	}
	return nil
}
