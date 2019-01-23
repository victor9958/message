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
	beego.Router("/admin/member-add", &controllers.MemberController{},"get:AddPage")//跳页面
	beego.Router("/admin/member-add-admin", &controllers.MemberController{},"post:AddAdmin")//跳页面
	beego.Router("/admin/member-list-data", &controllers.MemberController{},"get:ListData")
	beego.Router("/admin/member-del", &controllers.MemberController{},"delete:Del")
	beego.Router("/admin/change-job-page", &controllers.MemberController{},"get:ChangeJobPage")
	beego.Router("/admin/change-job", &controllers.MemberController{},"post:ChangeJob")


	//job
	beego.Router("/admin/job-list", &controllers.JobController{},"get:List")
	beego.Router("/admin/job-list-data", &controllers.JobController{},"get:ListData")
	beego.Router("/admin/job-del", &controllers.JobController{},"delete:Del")
	beego.Router("/admin/job-add", &controllers.JobController{},"get:AddPage")
	beego.Router("/admin/change-job-role-page", &controllers.JobController{},"get:ChangeJobRolePage")
	beego.Router("/admin/job-add-data", &controllers.JobController{},"post:Add")
	beego.Router("/admin/job-change-role", &controllers.JobController{},"post:ChangeRole")


	beego.Router("/ceshi", &controllers.BaseController{},"get:Ceshi")




}

