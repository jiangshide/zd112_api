package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int64
	Name       string
	Password   string
	Salt       string
	Status     int
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *User) TableName() string {
	return TableName("permission_user")
}

func (this *User) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *User) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *User) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *User) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *User) List(pageSize, offSet int) (list []*User, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
