package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"reflect"
	"github.com/jiangshide/GoComm/utils"
)

func init() {
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPsw := beego.AppConfig.String("db.psw")
	dbName := beego.AppConfig.String("db.name")
	timeZone := beego.AppConfig.String("db.timezone")
	maxConn, _ := beego.AppConfig.Int("maxConn")
	maxIdle, _ := beego.AppConfig.Int("maxIdle")
	if dbPort == "" {
		dbPort = "3306"
	}
	dns := dbUser + ":" + dbPsw + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	if timeZone != "" {
		dns += "&loc=" + url.QueryEscape(timeZone)
	}
	orm.RegisterDataBase("default", "mysql", dns, maxConn, maxIdle)
	orm.RegisterModel(new(User), new(UserProfile), new(UserLocation),new(Banner))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}

func Field(model interface{}) (fieldName string, fieldValue interface{}) {
	if model == nil {
		return fieldName, fieldValue
	}
	var field reflect.Type
	var value reflect.Value
	field = reflect.TypeOf(model).Elem()
	value = reflect.ValueOf(model).Elem()
	size := field.NumField()

	for i := 0; i < size; i++ {
		v := value.Field(i)
		fieldName = utils.StrFirstToLower(field.Field(i).Name)
		switch value.Field(i).Kind() {
		case reflect.Bool:
			fieldValue = v.Bool()
		case reflect.Int:
			if v.Int() != 0 {
				fieldValue = v.Int()
			}
		case reflect.Int64:
			if v.Int() != 0 {
				fieldValue = v.Int()
			}
		case reflect.String:
			fieldValue = v.String()
		}
		if fieldValue != nil && fieldValue != 0 {
			break
		}
	}
	return
}
