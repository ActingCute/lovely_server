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
	this.NeedPostData(count.Url)

	code := helper.SUCCESS

	err := models.AddCount(&count)

	if err != nil {
		code = helper.FAILED
	}
	this.SetReturnData(code, "love you", err)
}

// @Title Add And Get Count
// @Description 获取浏览数和添加浏览数
// @Param	body		body 	models.Count	true	"body for Count content"
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @Failure 403 body is empty
// @router /add_get [post]
func (this *CountController) AddAndGet() {
	var count, page_count models.Count
	var err error

	this.GetPostDataNotStop(&count)
	this.NeedPostData(count.Url)

	code := helper.SUCCESS

	//若url==/则是首页，只需增加网站浏览数即可
	if count.Url != "/" {
		page_count.Url = count.Url
		err = models.AddCount(&page_count)
		if err == nil {
			count.Url = "/"
			err = models.AddCount(&count)
		}
	} else {
		err = models.AddCount(&count)
	}

	if err != nil {
		code = helper.FAILED
	}

	data := make(map[string]interface{})
	data["web_count"] = count
	data["page_count"] = page_count

	this.SetReturnData(code, "love you", data)
}
