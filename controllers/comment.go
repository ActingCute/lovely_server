package controllers

import (
	"lovely_server/helper"
	"lovely_server/models"
	"time"
)

// Operations about Comment
type CommentController struct {
	BaseController
}

// @Title Add Comment
// @Description Add Comment
// @Param	body		body 	models.Comment	true	"body for Comment content"
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @Failure 403 body is empty
// @router /add [post]
func (this *CommentController) Add() {
	var comment models.Comment

	this.GetPostDataNotStop(&comment)

	this.NeedPostData(comment.UserId, comment.Url, comment.Content)

	code := helper.SUCCESS

	comment.UpdateTime = time.Now()

	err := models.AddComment(comment)

	if err != nil {
		code = helper.FAILED
	}
	this.SetReturnData(code, "love you", err)
}
