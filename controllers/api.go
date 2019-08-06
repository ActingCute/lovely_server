package controllers

import (
	"lovely_server/helper"
	"lovely_server/models"
)

// Operations about Api
type ApiController struct {
	BaseController
}

func (this *ApiController) GetRunTime() {
	this.SetReturnData(helper.SUCCESS, "love you", models.App.RunTime)
}
