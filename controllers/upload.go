package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"fmt"
	"github.com/jiangshide/GoComm/utils"
	"mime/multipart"
	"strings"
)

type UploadController struct {
	BaseController
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} upload success
// @Failure 403 user not exist
// @router / [get]
func (this *UploadController) Get() {
	this.response("success1")
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UploadController) Post() {
	f, fh, err := this.GetFile("file")
	beego.Info("--------f:", f, " | fh:", fh, " | err:", err)
	defer f.Close()
	fileName := fh.Filename
	sufix := "default"
	if strings.Contains(fh.Filename, ".") {
		sufix = fileName[strings.LastIndex(fileName, ".")+1:]
	}
	fileName = utils.Md5(this.userName+time.RubyDate+utils.GetRandomString(10)) + "_" + fmt.Sprint(time.Now().Unix()) + "." + sufix
	toFilePath := this.upload + sufix + "/" + fileName
	beego.Info("----------upload:", this.upload, " | sufix:", sufix, " | fileName:", fileName, " | toFilePath:", toFilePath)
	var size int64
	path := utils.GetCurrentDir(toFilePath)
	beego.Info("-----------path:", path)
	if err = this.SaveToFile("file", path); err == nil {
		beego.Info("------err:", err)
		this.response(size)

	} else {
		beego.Info("------err:", err)
		this.false(-1, err)
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /multipart [post]
func (this *UploadController) Multipart() {

	var allFiles map[string][]*multipart.FileHeader = this.Ctx.Request.MultipartForm.File
	beego.Info("------------allFiles:", allFiles)
	var keys []string
	var files []*multipart.FileHeader
	for k, v := range allFiles {
		keys = append(keys, k)
		files = append(files, v...)
	}
	for k, v := range files {
		f, err := v.Open()
		defer f.Close()
		if err != nil {
			continue
		}
		beego.Info("-------err:", err)
		fileName := v.Filename
		sufix := "default"
		if strings.Contains(v.Filename, ".") {
			sufix = fileName[strings.LastIndex(fileName, ".")+1:]
		}
		fileName = utils.Md5(this.userName+time.RubyDate+utils.GetRandomString(10)) + "_" + fmt.Sprint(time.Now().Unix()) + "." + sufix
		toFilePath := this.upload + sufix + "/" + fileName
		path := utils.GetCurrentDir(toFilePath)
		if err = this.SaveToFile(keys[k], path); err == nil {
			this.response(v.Size)
		} else {
			this.false(-1, err)
		}
	}
}
