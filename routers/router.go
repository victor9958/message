package routers

import (
	"github.com/astaxie/beego"
	"message/controllers"
	"message/filters"
)

func init() {
	//登陆  and   注册
	beego.Router("/login-page", &controllers.LoginController{},"get:LoginPage")
	beego.Router("/login", &controllers.LoginController{},"post:Login")


	//后台路由
	beego.InsertFilter("/admin/*",beego.BeforeRouter,filters.IsLogin)

	beego.Router("/admin/index", &controllers.IndexController{},"get:Index")
	beego.Router("/admin/wel", &controllers.BaseController{},"get:Wel")

	beego.Router("/admin/member-list", &controllers.MemberController{},"get:List")
	//beego.Router("/admin/member-list2", &controllers.MemberController{},"get:List2")
	//beego.Router("/admin/member-list3", &controllers.MemberController{},"get:List3")
	beego.Router("/admin/member-add", &controllers.MemberController{},"get:Add")//跳页面
	beego.Router("/admin/member-add-admin", &controllers.MemberController{},"post:AddAdmin")//跳页面
	beego.Router("/admin/member-list2-data", &controllers.MemberController{},"get:ListData")
	beego.Router("/admin/member-del", &controllers.MemberController{},"delete:Del")



	beego.Router("/ceshi", &controllers.BaseController{},"get:Ceshi")




}

