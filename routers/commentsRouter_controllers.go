package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["lovely_server/controllers:AdminController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:AdminController"],
        beego.ControllerComments{
            Method: "CheckLogin",
            Router: `/check_login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lovely_server/controllers:AdminController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lovely_server/controllers:CommentController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lovely_server/controllers:CommentController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lovely_server/controllers:CommentController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lovely_server/controllers:CountController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:CountController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lovely_server/controllers:CountController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:CountController"],
        beego.ControllerComments{
            Method: "AddAndGet",
            Router: `/add_get`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lovely_server/controllers:CountController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:CountController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lovely_server/controllers:QiniuController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:QiniuController"],
        beego.ControllerComments{
            Method: "UptokenKey",
            Router: `/uptoken_key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lovely_server/controllers:QiniuController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:QiniuController"],
        beego.ControllerComments{
            Method: "UptokenWeb",
            Router: `/uptoken_web`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lovely_server/controllers:TwitterController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:TwitterController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lovely_server/controllers:TwitterController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:TwitterController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lovely_server/controllers:TwitterController"] = append(beego.GlobalControllerRouter["lovely_server/controllers:TwitterController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
