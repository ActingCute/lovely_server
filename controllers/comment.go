package controllers

import (
	"lovely_server/models"
	"github.com/astaxie/beego"
	"lovely_server/helper"
)

// Operations about Comment
type CommentController struct {
	beego.Controller
}

// @Title Add Comment
// @Description Add Comment
// @Param	body		body 	models.Comment	true	"body for Comment content"
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @Failure 403 body is empty
// @router /add [post]
func (this *ApiController) Add() {
	var comment models.Comment

	this.NeedPostData(comment.UserId,comment.Url,comment.FatherId,comment.SonId)

	this.GetPostDataNotStop(&comment)
	
	code := helper.SUCCESS

	err := models.AddComment(comment)
	
	if err != nil {
		code = helper.FAILED
	}
	this.SetReturnData(code,"love you",nil)
}