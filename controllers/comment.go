package controllers

import (
	helper "lovely_server/helper"
	"lovely_server/models"
	"time"
)

// Operations about Comment
type CommentController struct {
	BaseController
}

// @Title Get Comment
// @Description Get Post Comment
// @Param	body		body 	models.Comment	true	"body for Comment content"
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @Failure 403 body is empty
// @router /get [post]
func (this *CommentController) Get() {
	var comment models.Comment

	this.GetPostDataNotStop(&comment)
	this.NeedPostData(comment.Url)

	code := helper.SUCCESS

	comments, err := models.GetCommentsByUrl(comment.Url)

	if helper.Error(err) {
		code = helper.FAILED
	}
	this.SetReturnData(code, "love you", comments)
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
	var user models.User
	//先获取用户数据

	this.GetPostDataNotStop(&user)

	this.NeedPostData(user.Name)

	err := models.AddUser(&user)

	if helper.Error(err) {
		this.SetReturnData(helper.FAILED, "love you", err)
		return
	}

	this.GetPostDataNotStop(&comment)

	this.NeedPostData(comment.Url, comment.Content)

	comment.UpdateTime = time.Now()
	comment.UserId = user.Id

	err = models.AddComment(&comment)

	code := helper.SUCCESS
	if err != nil {
		code = helper.FAILED
	} else {
		data := make(map[string]interface{})
		data["user"] = user
		data["comment"] = comment
		this.SetReturnData(code, "love you", data)
		return
	}
	this.SetReturnData(code, "love you", err)
}

// @Title Delete Comment
// @Description Delete Comment
// @Param	body		body 	models.Comment	true	"body for Comment content"
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @Failure 403 body is empty
// @router /delete [post]
func (this *CommentController) Delete() {

	this.NeedAdminLogin()

	var comment models.Comment

	this.GetPostDataNotStop(&comment)
	this.NeedPostData(comment.Id)

	err := models.DeleteComment(&comment)

	if err != nil {
		this.SetReturnData(helper.FAILED, err.Error(), nil)
	}
	this.SetReturnData(helper.SUCCESS, "love you", nil)
}
