package controllers

import (
	"strings"
	"time"

	"github.com/qiniu/api.v7/auth/qbox"

	"github.com/qiniu/api.v7/storage"

	"lovely_server/helper"
	"lovely_server/models"
)

// 七牛文档: https://developer.qiniu.com/kodo/sdk/1238/go
//
// 鉴权
// 七牛 Go SDK 的所有的功能，都需要合法的授权。授权凭证的签算需要七牛账号下的一对有效的Access Key和Secret Key，这对密钥可以通过如下步骤获得：
//
// 1. 开通七牛开发者帐号
// 2. 如果已有账号，直接登录七牛开发者后台，查看 Access Key 和 Secret Key
//
// Bucket : 新建存储空间，就是你的空间名
//
// 客户端上传凭证
// 客户端（移动端或者Web端）上传文件的时候，需要从客户自己的业务服务器获取上传凭证，而这些上传凭证是通过服务端的SDK来生成的，然后通过客户自己的业务API分发给客户端使用。
// 根据上传的业务需求不同，七牛云 Go SDK支持丰富的上传凭证生成方式。
//
// import (
//     "github.com/qiniu/api.v7/auth/qbox"
//     "github.com/qiniu/api.v7/storage"
// )
// accessKey := "your access key"
// secretKey := "your secret key"
// mac := qbox.NewMac(accessKey, secretKey)

// 简单上传的凭证
// 最简单的上传凭证只需要AccessKey，SecretKey和Bucket就可以
// bucket:="your bucket name"
// putPolicy := storage.PutPolicy{
//         Scope: bucket,
// }
// putPolicy.Expires = 7200 //示例2小时有效期
// mac := qbox.NewMac(accessKey, secretKey)
// upToken := putPolicy.UploadToken(mac)

var QINIU_QBOX_MAC *qbox.Mac

type QiniuController struct {
	BaseController
}

func init() {
	QINIU_QBOX_MAC = qbox.NewMac(models.QiNiu.AccessKey, models.QiNiu.SecretKey)
}

// @Title Get Uptoken Web
// @Description Get Uptoken Web
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @Failure 403 body is empty
// @router /uptoken_web [get]
func (this *QiniuController) UptokenWeb() {
	upToken := make(map[string]string)
	upToken["token"] = Uptoken(models.QiNiu.Space)
	upToken["domain"] = models.QiNiu.Domain
	this.SetReturnData(helper.SUCCESS, "ok", upToken)
}

// @Title Get Uptoken Key
// @Description Get Uptoken Key
// @Success 10000 {struct} helper.RestfulReturn
// @Failure 10001 {struct} helper.RestfulReturn
// @Failure 403 body is empty
// @router /uptoken_key [get]
func (this *QiniuController) UptokenKey() {
	resultData := make(map[string]interface{})
	resultData["domain"] = models.QiNiu.Domain
	resultData["app_domain"] = models.App.ApiUrl
	resultData["qiniu_enable"] = models.QiNiu.Enable
	fileName := this.GetString("file_name")
	//若不存在文件名，则用随机数，否则多文件会出问题
	fileName = strings.Replace(fileName, "\"", "", -1)
	if len(fileName) < 1 {
		fileName = helper.GetToken(10, helper.KC_RAND_KIND_ALL)
	}

	timeStr := time.Now().Format("2006-01-02 15:04:05")
	timeStr = strings.Replace(timeStr, " ", "_", -1)
	timeStr = strings.Replace(timeStr, ":", "-", -1)

	key := "qiniu/upload/" + timeStr + "/" + fileName

	//先判断是否开启了七牛上传，否就上传到本地咯
	// if !models.QiNiu.Enable {
	// 	resultData["key"] = "upload/" + timeStr + "/" + fileName
	// 	resultData["token"] = "" // 3600是一小时。
	// 	resultData["action"] = models.App.ApiUrl + "/v1/api/upload_file"
	// 	this.SetReturnData(SUCCESS, "ok", resultData)
	// 	this.Finish()
	// 	return
	// }

	resultData["key"] = key
	resultData["token"] = UptokenWithKey(models.QiNiu.Space, key, 3600*24) // 3600是一小时。
	resultData["action"] = models.QiNiu.Action

	this.SetReturnData(helper.SUCCESS, "ok", resultData)
	this.Finish()
}

func Uptoken(bucketName string) string {
	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	putPolicy := storage.PutPolicy{
		Scope: bucketName,
	}

	putPolicy.Expires = 60 * 60 * 24 * 3 // token过期时间设置为3天.

	return putPolicy.UploadToken(QINIU_QBOX_MAC)
}

func UptokenWithKey(bucket, key string, expires uint32) string {
	putPolicy := storage.PutPolicy{
		Scope: bucket + ":" + key,
	}

	putPolicy.Expires = 60 * 60 * 24 // token过期时间设置为1天.

	return putPolicy.UploadToken(QINIU_QBOX_MAC)
}
