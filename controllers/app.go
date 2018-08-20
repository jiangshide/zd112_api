package controllers

import (
	"zd112_api/models"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type AppController struct {
	BaseController
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router / [get]
func (this *AppController) Get() {
	channel := new(models.Channel)
	result, _ := channel.List(this.Page())
	this.response(result)
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /update [get]
func (this *AppController) Update() {
	update := new(models.Update)
	update.Name = this.getString("name", DEFAULT_ISNULL, 1)
	update.Channel = this.getString("channel", DEFAULT_ISNULL, 1)
	update.Platform = this.getInt("platform", 0)
	update.Version = this.getString("version", DEFAULT_ISNULL, 1)
	var maps []orm.Params
	if _, err := update.Query().Filter("name", update.Name).Filter("channel", update.Channel).Filter("platform", update.Platform).Filter("version__gt", update.Version).Values(&maps, "status", "content", "url"); err != nil {
		this.false(-1, err)
	}
	for _, v := range maps {
		v["Url"] = this.host + fmt.Sprint(v["Url"])
	}
	update.Url = this.host + update.Url
	this.response(maps)
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /stop [get]
func (this *AppController) Stop() {
	stop := new(models.Stop)
	var maps []orm.Params
	_, err := stop.Query().Values(&maps,   "url")
	if err != nil {
		this.false(-1,err)
	}
	if len(maps) == 0{
		this.false(-1,"no data!")
	}
	this.response(maps[0])
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /advert [get]
func (this *AppController) Advert() {
	advert := new(models.Advert)
	advert.Type = this.getInt("type", 2)
	var maps []orm.Params
	_, err := advert.Query().Filter("type", advert.Type).Values(&maps, "name", "url")
	if err != nil {
		this.false(-1, err)
	}
	if len(maps) == 0{
		this.false(-1,"no data!")
	}
	for _, v := range maps {
		v["Url"] = this.host + fmt.Sprint(v["Url"])
		break
	}
	this.response(maps[0])
}
