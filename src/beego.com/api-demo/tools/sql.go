package tools

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

// 拼接参数，形成分页查询的 sql 语句
func GenerateListSQL(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) string {
	sql := "select "
	if len(fields) == 0 {
		sql += "* "
	} else {
		for i, val := range fields {
			if i != 0 {
				sql += ", "
			}
			sql += val + " "
		}
	}
	sql += "from gateway where "
	for k, v := range query {
		sql += k + " = '" + v + "' and "
	}
	if len(query) > 0 {
		sql = beego.Substr(sql, 0, len(sql)-4)
	}

	sql += "order by "
	for i, val := range sortby {
		if i != 0 {
			sql += ", "
		}
		sql += val + " "
		if len(order)-1 > i {
			sql += order[i] + " "
		}
	}
	sql += "limit " + strconv.FormatInt(offset, 10) + ", " + strconv.FormatInt(limit, 10)

	logs.Info("GenerateListSQL result: " + sql)

	return sql
}
