package controllers

import (
	"zd112_api/models"
	"zd112/utils"
	"time"
)

// Users APi
type UserController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserController) Reg() {
	user := new(models.User)
	user.LoginName = this.getString("username", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	password := this.getString("password", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	rePassword := this.getString("repassword", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	if password != rePassword {
		this.false(PASSWORD_DIFFER, nil)
	}
	user.Salt = utils.GetRandomString(10)
	user.Password = utils.Md5(password + user.Salt)
	user.CreateTime = time.Now().Unix()
	if _, err := user.Add(); err != nil {
		this.false(DB_INSERT_FALSE, err)
	}
	this.response(user)
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (this *UserController) Login() {
	this.checkToken()
	user := new(models.User)
	user.LoginName = this.getString("username", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	password := this.getString("password", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	if err := user.Query(); err != nil {
		this.false(DB_QUERY_FALSE, err)
	}
	if user.Status == 0 {
		this.false(USER_NOT_ACTIVE, nil)
	} else if user.Status == 2 {
		this.false(USER_FORBIDDEN, nil)
	} else if user.Password != utils.Md5(password+user.Salt) {
		this.false(USER_PASSWORD_ERR, nil)
	}
	user.UpdateTime = time.Now().Unix()
	if _, err := user.Update(); err != nil {
		this.false(DB_UPDATE_FALSE, err)
	}
	this.response(this.token())
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (this *UserController) Logout() {
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *UserController) GetAll() {
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /:userId [get]
func (this *UserController) Get() {
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (this *UserController) Delete() {
	user := new(models.User)
	user.Id = this.getInt64("userId", 0)
	if _, err := user.Del(); err != nil {
		this.false(DB_DELETE_FALSE, err)
	}
	this.Delete()
}
