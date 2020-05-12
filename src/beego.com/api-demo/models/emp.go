package models

import (
	"fmt"
	_ "github.com/astaxie/beego/config/yaml"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/go-sql-driver/mysql"
)

type Emp struct {
	Id    int32 `orm:"column(empno);pk"`
	Ename string
	Job   string
	Mgr   int32
	// todo
	Hiredate string
	Sal      float64
	Comm     float64
	Dept     *Dept `orm:"rel(fk);column(deptno)" valid:"Range(10, 40)"`
}

type Dept struct {
	Id    int16 `orm:"column(deptno);pk"`
	Dname string
	Loc   string
	Emps  []*Emp `orm:"reverse(many)"`
}

func init() {
	// 如果要让 beego 实现将sql执行结果映射成结构体时，就需要先对结构体进行注册
	orm.RegisterModel(new(Emp), new(Dept))
}

func QueryAllBySQL() (emps []*Emp, err error) {
	sql := "select * from emp"
	count, err := orm.NewOrm().Raw(sql).QueryRows(&emps)
	logs.Debug("%s result count: %d", sql, count)
	return
}

func QueryValues() (datas []orm.Params, err error) {
	// 查询出部门平均薪水高于1w的
	sql := "select t.deptno, avg(t.sal) avg_sal from ( select * from emp join dept using(deptno)) t GROUP BY t.deptno having avg_sal > 10000;"

	// Values 将结果转换成 []map[string]interface 类型
	count, err := orm.NewOrm().Raw(sql).Values(&datas)
	logs.Debug("%s result count: %d", sql, count)
	return
}

func QueryBuilder() (sql string, err error) {
	queryBuilder, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		return "", err
	}
	// queryBuilder 帮助开发者形成复杂的 sql 语句
	queryBuilder.
		Select("emp.ename", "dept.dname").
		From("emp").
		InnerJoin("dept").
		On("emp.deptno = dept.deptno").
		Where("sal > 2400").
		OrderBy("sal").Desc().
		Limit(2).Offset(1)
	sql = queryBuilder.String()
	return
}

func QuerySeter() (*[]orm.Params, error) {
	querySeter := orm.NewOrm().QueryTable("emp")
	if querySeter == nil {
		return nil, fmt.Errorf("querySeter is nil")
	}
	var data []orm.Params
	_, err := querySeter.
		Filter("sal__gt", 2500).
		Filter("ename__icontains", "O").
		Distinct().
		OrderBy("empno", "-deptno").
		Filter("deptno__exact", 30).
		// SetCond 自定义条件表达式，与 Filter 同时使用时，Filter会失效
		//SetCond(orm.NewCondition().AndNot("comm__isnull", true)).
		Limit(5).
		Values(&data)
	return &data, err
}

func Validate(emp *Emp) (errors []error) {
	valid := validation.Validation{}
	ok, err := valid.Valid(emp)
	if err != nil {
		errors = append(errors, err)
	}
	if !ok {
		logs.Debug("validate error count : %d", len(valid.Errors))
		for _, err = range valid.Errors {
			errors = append(errors, err)
		}
	}
	return
}

// 根据 depe id，查出dept，并关联查询出该dept对应的多个emps
func OneToMany() (*Dept, error) {
	dept := &Dept{Id: 20}
	ormer := orm.NewOrm()
	// 读取 id=20 的 dept 信息
	err := ormer.Read(dept)
	if err != nil {
		return nil, err
	}
	// 将 dept 中的 emps 字段进行表关联查询
	_, err = ormer.LoadRelated(dept, "Emps")
	return dept, err
}

// 查询 emp 表并与 dept 做内连接，对结果支持条件查询
func ManyToOne() ([]*Emp, error) {
	var emps []*Emp
	ormer := orm.NewOrm()
	// RelatedSel() 实现自动做关联查询
	_, err := ormer.QueryTable("emp").
		/*Filter("dept__dname", "RESEARCH").*/ RelatedSel().All(&emps)
	return emps, err
}
