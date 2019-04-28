package controllers

import (
	"lovely_server/helper"
	"lovely_server/models"
)

// Operations about count
type CountController struct {
	BaseController
}

// @Title Add Comment
// @Description Add Post Count
// @Param	body		body 	models.Comment	true	"body for Comment content"
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
