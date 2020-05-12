package tools

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type DatabaseCheck struct {
}

func (dc *DatabaseCheck) Check() error {
	if dc.isConnected() {
		return nil
	} else {
		return errors.New("database connection is not ready")
	}
}

func (dc *DatabaseCheck) isConnected() bool {
	db, err := orm.GetDB()
	if err != nil {
		logs.Error("orm.GetDB error :%v", err)
	}
	err = db.Ping()
	if err == nil {
		return true
	}
	return false
}
