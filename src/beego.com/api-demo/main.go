package main

import (
	_ "beego.com/api-demo/models"
	_ "beego.com/api-demo/routers"
	"beego.com/api-demo/tools"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego/toolbox"
	"strconv"
	"time"
)

func init() {
	// 设置 yaml 格式配置文件
	err := beego.LoadAppConfig("yaml", "conf/app.yaml")
	if err != nil {
		println(err.Error())
		return
	}

	// 设置日志输出显示行号，默认不配置就会输出
	logs.SetLogFuncCall(true)
	// orm 开启 Debug 日志等级
	orm.Debug = true

	// 注册 mysql 驱动
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logs.Error("orm.RegisterDriver error: %v", err)
		return
	}
	// 注册默认的数据源
	port, err := beego.AppConfig.Int64("MysqlPort")
	if err != nil {
		logs.Error("get MysqlPort property error: %v", err)
		return
	}
	dataSource := beego.AppConfig.String("MysqlUn") + ":" + beego.AppConfig.String("MysqlPwd") +
		"@tcp(" + beego.AppConfig.String("MysqlHost") + ":" + strconv.FormatInt(port, 10) + ")/demo?charset=utf8"
	logs.Info("dataSource: ", dataSource)
	maxIdle := 30
	maxConn := 30
	err = orm.RegisterDataBase("default", "mysql", dataSource, maxIdle, maxConn)
	if err != nil {
		logs.Error("orm.RegisterDataBase error :%v", err)
		return
	}
	db, err := orm.GetDB()
	if err != nil {
		logs.Error("orm.GetDB error :%v", err)
	}
	db.SetConnMaxLifetime(time.Minute * 60)

	// 自动根据 struts 进行建表
	// orm.RunSyncdb("default", true, false)

	// 添加对 db 的健康检查
	toolbox.AddHealthCheck("default", &tools.DatabaseCheck{})

	// 为 session 存储引擎设置数据源
	beego.BConfig.WebConfig.Session.SessionProviderConfig = dataSource
}

func main() {
	// 开启测试定时任务
	startTask()
	defer toolbox.StopTask()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func startTask() {
	task := createTestTask()
	toolbox.AddTask(task.Taskname, task)
	toolbox.StartTask()
}

func createTestTask() *toolbox.Task {
	testFunc := func() error {
		logs.Debug("this is a test task")
		return nil
	}
	return toolbox.NewTask("test", "0 */10 * * * *", testFunc)
}
