package controllers

import (
	helper "lovely_server/helper"
	"lovely_server/models"
)

// Operations about Admin
type AdminController struct {
	BaseController
}

// @Title Admin Login
// @Description Admin Login
// @Param	body		body 	models.Admin	true	"body for Admin content"
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @Failure 403 body is empty
// @router /login [post]
func (this *AdminController) Login() {
	var admin models.Admin

	this.GetPostDataNotStop(&admin)
	this.NeedPostData(admin.UserName, admin.PassWord)

	has, admin, err := models.CkeckAdmin(admin)

	if helper.Error(err) {
		this.SetReturnData(helper.FAILED, "login fail", err.Error())
		return
	}
	if !has {
		this.SetReturnData(helper.FAILED, "pass_word error", nil)
		return
	}
	this.SetReturnData(helper.SUCCESS, "login ok", admin)
	return
}

// @Title Check Admin Login
// @Description Check Admin Login
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @router /check_login [get]
func (this *AdminController) CheckLogin() {
	var admin models.Admin

	this.GetPostDataNotStop(&admin)
	this.NeedPostData(admin.UserName, admin.PassWord)
	admin = this.NeedAdminLogin()

	this.SetReturnData(helper.SUCCESS, "login ok", admin)
}
