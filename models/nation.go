package models

import "github.com/astaxie/beego/orm"

type Nation struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	CreateId   int64  `json:"create_id"`
	UpdateId   int64  `json:"update_id"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func (this *Nation) TableName() string {
	return TableName("nation")
}

func (this *Nation) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Nation) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Nation) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Nation) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Nation) QueryAll() (list []*Nation, err error) {
	orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).All(&list)
	return
}

func (this *Nation) List(pageSize, offSet int) (list []*Nation, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
