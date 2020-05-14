package models

import (
	"beego.com/api-demo/tools"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type Gateway_ struct {
	Namespace         string `json:"namespace,omitempty"`
	Name              string `json:"name,omitempty"`
	Json              string `json:"json,omitempty"`
	CreationTimestamp string `json:"create_time,omitempty"`
}

type Gateway struct {
	Metadata Metadata    `json:"metadata,omitempty"`
	Spec     GatewaySpec `json:"spec,omitempty"`
}

type Metadata struct {
	Labels            map[string]string `json:"labels,omitempty"`
	Name              string            `json:"name,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	CreationTimestamp Time              `json:"creationTimestamp,omitempty"`
}

type Time struct {
	time.Time
}

type GatewaySpec struct {
	Servers  []*Server         `json:"servers,omitempty"`
	Selector map[string]string `json:"selector,omitempty"`
}

type Server struct {
	Port  *Port              `json:"port,omitempty"`
	Hosts []string           `json:"hosts,omitempty"`
	Tls   *ServerTLSSettings `json:"tls,omitempty"`
}

type ServerTLSSettings struct {
	ServerCertificate string `json:"server_certificate,omitempty"`
	PrivateKey        string `json:"private_key,omitempty"`
}

type Port struct {
	Number   uint32 `json:"number,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Name     string `json:"name,omitempty"`
}

func AddGateway(gateway *Gateway) error {
	parse, _ := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	gateway.Metadata.CreationTimestamp.Time = parse
	bytes, err := json.Marshal(gateway)
	if err != nil {
		return err
	}
	gateway_ := Gateway_{
		Namespace: gateway.Metadata.Namespace,
		Name:      gateway.Metadata.Name,
		Json:      string(bytes),
	}
	sql := "insert into gateway(namespace, name, json, create_time) values(?, ?, ?, ?)"
	result, err := orm.NewOrm().Raw(sql, gateway_.Namespace, gateway_.Name, gateway_.Json, time.Now().UTC().Format(time.RFC3339)).Exec()
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	logs.Info("method 'AddGateway' rows affected count: %d", affected)
	return nil
}

func GetGatewayById(int64) (*Gateway, error) {

	return nil, nil
}

// select fileds from table_name where query order by sortby order limit limit, offset
func GetAllGateway(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (gateways []*Gateway_, err error) {
	sql := tools.GenerateListSQL(query, fields, sortby, order, offset, limit)
	// todo 无需转成 Gateway_，直接返回map。此处还需要形成端口配置并返回
	var lst []orm.Params
	_, err = orm.NewOrm().Raw(sql).Values(&lst)
	bytes, err := json.Marshal(lst)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &gateways)
	return
}

func DeleteGateway(ns, name string) error {
	sql := "delete from gateway where namespace = ? and name = ?"
	result, err := orm.NewOrm().Raw(sql, ns, name).Exec()
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	logs.Info("method 'DeleteGateway' rows affected count: %d", affected)
	return nil
}
