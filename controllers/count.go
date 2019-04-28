package controllers

import (
	helper "lovely_server/helper"
	"lovely_server/models"
)

// Operations about count
type CountController struct {
	BaseController
}

// @Title Get Count
// @Description Get Post Count
// @Param	body		body 	models.Count	true	"body for Count content"
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @Failure 403 body is empty
// @router /get [post]
func (this *CountController) Get() {
	var count models.Count

	this.GetPostDataNotStop(&count)
	this.NeedPostData(count.Url)

	code := helper.SUCCESS

	counts, err := models.GetCountsByUrl(count.Url)

	if err != nil {
		code = helper.FAILED
	}
	this.SetReturnData(code, "love you", counts)
}

// @Title Add Count
// @Description Add Post Count
// @Param	body		body 	models.Count	true	"body for Count content"
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @Failure 403 body is empty
// @router /add [post]
func (this *CountController) Add() {
	var count models.Count

	this.GetPostDataNotStop(&count)
	this.NeedPostData(count.Url, count.Count)

	code := helper.SUCCESS

	err := models.AddCount(count)

	if err != nil {
		code = helper.FAILED
	}
	this.SetReturnData(code, "love you", err)
}
