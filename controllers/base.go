package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	helper "lovely_server/helper"
	"lovely_server/models"
	"strings"
)

type BaseController struct {
	beego.Controller
}

type Page struct {
	Per int `json:"per"`
	Cur int `json:"cur"`
}

func (this *BaseController) Prepare() {
	helper.Debug("好无聊喔~", models.Domain)
	this.NeedSQLConnentSuccess()
}

func (this *BaseController) Options() {
	this.AllowCross() //允许跨域
	this.SetReturnData(helper.SUCCESS, "ok", nil)
}

//数据库要连接正常
func (this *BaseController) NeedSQLConnentSuccess() {
	if !models.DBOk {
		helper.Error("数据库连接失败 ----")
		this.SetReturnData(helper.SQL_CONNECT_ERROR, "db connent fail", nil)
		this.Finish()
		this.StopRun()
	}
}

//跨域
func (this *BaseController) AllowCross() {
	beego.Debug("AllowCross --- ", models.Domain)
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", models.Domain)                 //允许访问源
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")    //允许post访问
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	this.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
}

//返回json数据
func (this *BaseController) SetReturnData(result helper.Status, message string, data interface{}) {
	rt := &helper.RestfulReturn{Result: result, Message: message, Data: data}
	this.Data["json"] = rt
	this.ServeJSON()
}

//非空数据
func (this *BaseController) NeedPostData(data ...interface{}) {
	for _, v := range data {
		switch t := v.(type) {
		case string:
			s := v.(string)
			if s == "" {
				this.SetReturnData(helper.PARAMETER_ERROR, "Lack of parameters 1", nil)
				this.Finish()
				this.StopRun()
			}
			if strings.Replace(s, " ", "", -1) == "" {
				this.SetReturnData(helper.PARAMETER_ERROR, "Lack of parameters 2", nil)
				this.Finish()
				this.StopRun()
			}
			break
		case int, int64: //判断不能为0
			if v.(int64) == 0 {
				this.SetReturnData(helper.PARAMETER_ERROR, "Lack of parameters 3", nil)
				this.Finish()
				this.StopRun()
			}
			break
		default:
			_ = t
			helper.Error("不知道啥")
			this.SetReturnData(helper.PARAMETER_ERROR, "Lack of parameters 4", nil)
			this.Finish()
			this.StopRun()
		}
	}
}

//获取post数据
func (this *BaseController) GetPostData(data interface{}) {
	err := json.Unmarshal(this.Ctx.Input.RequestBody, data)
	if helper.Error(err) {
		this.SetReturnData(helper.PARAMETER_ERROR, "Lack of parameters", nil)
		this.Finish()
		this.StopRun()
	}
}

func (this *BaseController) GetPostDataNotStop(data interface{}) {
	json.Unmarshal(this.Ctx.Input.RequestBody, data)
	helper.Debug(string(this.Ctx.Input.RequestBody))
}

func (this *BaseController) NeedAdminLogin() models.Admin {
	Token := this.Ctx.Input.Header("Authorization")
	helper.Debug("Token :", Token)
	if len(Token) != 32 {
		this.SetReturnData(helper.PARAMETER_ERROR, "no token", nil)
		this.Finish()
		this.StopRun()
	}
	admin := models.Admin{Token: Token}
	has, admin, err := models.GetAdminByToken(admin)
	if !has {
		this.SetReturnData(helper.FAILED, "登陆过期", nil)
		this.Finish()
		this.StopRun()
	}
	if helper.Error(err) {
		this.SetReturnData(helper.FAILED, err.Error(), err)
		this.Finish()
		this.StopRun()
	}
	return admin
}
