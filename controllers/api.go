package controllers

import (
	"lovely_server/helper"
	"lovely_server/models"
	"time"
)

// Operations about Api
type ApiController struct {
	BaseController
}

func (this *ApiController) GetRunTime() {
	this.SetReturnData(helper.SUCCESS, "love you", time.Now().Unix() - models.App.RunTime)
}
