package controllers

import (
	helper "lovely_server/helper"
	"lovely_server/models"
	"time"
)

// Operations about Twitter
type TwitterController struct {
	BaseController
}

// @Title Delete Twitter
// @Description Delete Twitter
// @Param	body		body 	models.Twitter	true	"body for Twitter content"
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @Failure 403 body is empty
// @router /delete [post]
func (this *TwitterController) Delete() {

	this.NeedAdminLogin()

	var Twitter models.Twitter

	this.GetPostDataNotStop(&Twitter)
	this.NeedPostData(Twitter.Id)

	err := models.DeleteTwitter(&Twitter)

	if err != nil {
		this.SetReturnData(helper.FAILED, err.Error(), nil)
	}
	this.SetReturnData(helper.SUCCESS, "love you", nil)
}


// @Title Add Twitter
// @Description Add Twitter
// @Param	body		body 	models.Twitter	true	"body for Twitter content"
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @Failure 403 body is empty
// @router /add [post]
func (this *TwitterController) Add() {

	this.NeedAdminLogin()

	var Twitter models.Twitter

	this.GetPostDataNotStop(&Twitter)
	this.NeedPostData(Twitter.Content)

	Twitter.UpdateTime = time.Now()

	err := models.AddTwitter(&Twitter)

	if err != nil {
		this.SetReturnData(helper.FAILED, err.Error(), nil)
	}
	this.SetReturnData(helper.SUCCESS, "love you", nil)
}