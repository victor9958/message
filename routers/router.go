package routers

import (
	"github.com/astaxie/beego"
	"message/controllers"
	"message/filters"
)

func init() {
	//登陆  and   注册
	beego.Router("/login-page", &controllers.LoginController{}, "get:LoginPage")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")


	beego.InsertFilter("/admin/*",beego.BeforeRouter,filters.IsLogin)
	//beego.InsertFilter("/admin/member-list",beego.BeforeRouter,filters.Auth)
	//beego.InsertFilter("/admin/list-jo",beego.BeforeRouter,filters.Auth)
	beego.InsertFilter("/admin/list*",beego.BeforeRouter,filters.Auth)
	//beego.InsertFilter("/*.-list",beego.BeforeRouter,filters.Auth)


	beego.Router("/admin/list-job", &controllers.JobController{}, "get:List")
	beego.Router("/admin/list-member", &controllers.MemberController{}, "get:List")

	beego.Router("/admin/index", &controllers.IndexController{},"get:Index")
	beego.Router("/admin/wel", &controllers.BaseController{},"get:Wel")

	beego.Router("/admin/member-add", &controllers.MemberController{},"get:AddPage")//跳页面
	beego.Router("/admin/member-add-admin", &controllers.MemberController{},"post:AddAdmin")//跳页面
	beego.Router("/admin/member-list-data", &controllers.MemberController{},"get:ListData")
	beego.Router("/admin/member-del", &controllers.MemberController{},"delete:Del")


	//job

	beego.Router("/admin/job-list-data", &controllers.JobController{}, "get:ListData")
	beego.Router("/admin/job-del", &controllers.JobController{}, "delete:Del")
	beego.Router("/admin/job-add", &controllers.JobController{}, "get:AddPage")
	beego.Router("/admin/job-edit-page", &controllers.JobController{}, "get:EditPage")
	beego.Router("/admin/job-edit-data", &controllers.JobController{}, "post:Edit")
	beego.Router("/admin/change-job-role-page", &controllers.JobController{}, "get:ChangeJobRolePage")
	beego.Router("/admin/job-add-data", &controllers.JobController{}, "post:Add")
	beego.Router("/admin/job-change-role", &controllers.JobController{}, "post:ChangeRole")

	//user
	beego.Router("/admin/list-user", &controllers.Usercontroller{}, "get:List")
	beego.Router("/admin/user-list-data", &controllers.Usercontroller{}, "get:ListData")
	beego.Router("/admin/user-add", &controllers.Usercontroller{}, "get:AddPage")
	beego.Router("/admin/change-user-add", &controllers.Usercontroller{}, "post:Add")
	beego.Router("/admin/change-user-edit", &controllers.Usercontroller{}, "post:Edit")
	beego.Router("/admin/change-user-edit-page", &controllers.Usercontroller{}, "get:EditPage")

	beego.Router("/ceshi", &controllers.BaseController{}, "get:Ceshi")

	auth := beego.NewNamespace("/admin",
	)

	beego.AddNamespace(auth)
}
