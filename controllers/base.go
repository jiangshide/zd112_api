package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/jiangshide/GoComm/utils"
	"time"
	"fmt"
)

type BaseController struct {
	beego.Controller
	version    string
	controller string
	action     string
	page       int
	pageSize   int
	offSet     int
	userName   string
	upload     string
	host       string

	code    int
	msg     interface{}
	content interface{}
}

func (this *BaseController) Prepare() {
	this.version = beego.AppConfig.String("version")
	controller, action := this.GetControllerAndAction()
	this.controller = strings.ToLower(controller[0 : len(controller)-10])
	this.action = strings.ToLower(action)
	this.page, _ = beego.AppConfig.Int("page")
	this.pageSize, _ = beego.AppConfig.Int("pageSize")
	this.offSet = (this.page - 1) * this.pageSize
	this.upload = this.GetString("upload", "/static/upload/")
	this.host = beego.AppConfig.String("host")
}

func (this *BaseController) header(key string) string {
	return this.Ctx.Request.Header.Get(key)
}

func (this *BaseController) checkToken() bool {
	auth := this.header(AUTH)
	if len(auth) == 0 {
		this.false(TOKEN_INVALIDATE, "auth is null!")
	}
	if _, err := utils.UnToken(auth, this.userName); err != nil {
		this.false(TOKEN_PRODUCE_FALSE, err)
	} else {
		return true
	}
	return false
}

func (this *BaseController) Page() (int, int) {
	num, _ := this.GetInt("page", 1)
	size, _ := this.GetInt("pageSize", 10)
	num = size * num
	return size, num
}

func (this *BaseController) response(content interface{}) {
	res := make(map[string]interface{}, 0)
	res["code"] = RES_OK
	//res["version"] = this.version
	res["date"] = time.Now().Unix()
	if content != nil {
		res["res"] = content
	}
	this.Data["json"] = res
	this.ServeJSON()
}

func (this *BaseController) token() string {
	tokenMap := make(map[string]interface{}, 0)
	tokenMap["code"] = this.version
	//tokenMap["version"] = this.version
	tokenMap["username"] = this.userName
	if token, err := utils.Token(tokenMap, this.userName, time.Now().Add(time.Hour * 1).Unix()); err != nil {
		this.false(TOKEN_PRODUCE_FALSE, nil)
		return ""
	} else {
		return token
	}
}

func (this *BaseController) false(code int, err interface{}) {
	res := make(map[string]interface{}, 0)
	res["code"] = code
	//res["version"] = this.version
	if err != nil {
		res["msg"] = fmt.Sprint(err)
	} else {
		res["msg"] = msg[code]
	}
	res["date"] = time.Now().Unix()
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) getInt(key string, defaultValue int) int {
	resInt, err := this.GetInt(key, defaultValue)
	if err != nil {
		this.false(GETINT_FALSE, err)
	}
	return resInt
}

func (this *BaseController) getInt64(key string, defaultValue int64) int64 {
	resInt64, err := this.GetInt64(key, defaultValue)
	if err != nil {
		this.false(GETINT64_FALSE, err)
	}
	return resInt64
}

func (this *BaseController) getString(key string, code, minSize int) string {
	value := strings.TrimSpace(this.GetString(key, ""))
	if len(value) == 0 {
		if code == DEFAULT_ISNULL {
			this.false(code, "你输入的"+key+"项不能为空!")
		} else {
			this.false(code, nil)
		}
	} else if len(value) < minSize {
		this.false(DEFAULT_SIZE_LOW, "你输入的"+key+"项的长度不能小于"+fmt.Sprint(minSize)+"个字符!")
	}
	return value
}

func (this *BaseController) getIp() string {
	return this.Ctx.Input.IP()
}
