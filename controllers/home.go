package controllers

import (
	"github.com/astaxie/beego"
	"zd112_api/models"
)

// Operations about Users
type HomeController struct {
	BaseController
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router / [get]
func (this *HomeController) Get() {
	beego.Info("---------home")
	banner := new(models.Banner)
	result,_:=banner.List(this.pageSize,this.offSet)

	this.response(result)
}