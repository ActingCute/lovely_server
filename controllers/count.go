package controllers

import (
	"lovely_server/models"
	"lovely_server/helper"
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
	var comment models.Comment

	this.NeedPostData(comment)

	this.GetPostDataNotStop(&comment)
	
	code := helper.SUCCESS

	err := models.AddComment(comment)
	
	if err != nil {
		code = helper.FAILED
	}
	this.SetReturnData(code,"love you",nil)
}