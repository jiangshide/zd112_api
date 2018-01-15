package controllers

const (
	AUTH             = "auth"
	DEFAULT_TIPS     = "该项不能为空!"
	DEFAULT_MIN_SIZE = 6
	//the base
	RES_OK              = 0
	FALSE               = 1
	TOKEN_PRODUCE_FALSE = 2
	TOKEN_INVALIDATE    = 3

	//the systen:10~300
	MAC_ISNULL     = 10
	DEVICE_ISNULL  = 11
	GETINT_FALSE   = 12
	GETINT64_FALSE = 13

	//the user:301~500
	DEFAULT_ISNULL    = 301
	DEFAULT_SIZE_LOW  = 302
	PASSWORD_DIFFER   = 303
	USER_NOT_ACTIVE   = 304
	USER_FORBIDDEN    = 305
	USER_PASSWORD_ERR = 306

	//the db:501~800
	DB_INSERT_FALSE = 501
	DB_DELETE_FALSE = 502
	DB_UPDATE_FALSE = 503
	DB_QUERY_FALSE  = 503
	//the net:801~1200
	//the other:1201~...
)

var msg = map[int]interface{}{
	//the system:10~300
	MAC_ISNULL:    "Mac地址不能为空!",
	DEVICE_ISNULL: "设备名称不能为空!",
	//the user:301~500
	PASSWORD_DIFFER:   "密码不一致!",
	USER_NOT_ACTIVE:   "该账号未激活!",
	USER_FORBIDDEN:    "该账号已被禁用!",
	USER_PASSWORD_ERR: "输入密码错误!",
	//the db:501~800

	//the net:801~1200

	//the other:1201~...
}

type ErrorController struct {
	BaseController
}

func (this *ErrorController) Error404() {
	this.Data["content"] = "正在开发中..."
	this.false(-1, "404错误")
}

func (this *ErrorController) Error501() {
	this.Data["content"] = "server error"
	this.false(-1, "501错误")
}

func (this *ErrorController) ErrorDb() {
	this.Data["content"] = "database is now down"
	this.false(-1, "db操作错误")
}
