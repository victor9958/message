package controllers

import (
	"github.com/astaxie/beego/orm"
	"message/funcs"
	"message/model"
	"strconv"
	"time"
)

type MemberController struct {
	BaseController
}

/*
	用户列表 ui框架
 */
func(this *MemberController)List(){
	this.TplName="member-list.html"
}

/*
	跳添加页面
 */
func(this *MemberController)AddPage(){
	this.TplName="member-add.html"
}

func(this *MemberController)AddAdmin(){
	var admin model.Admin
	name := this.GetString("name")
	if name == "" {
		this.ReturnJson(map[string]string{"message":"名称必填"},400)
	}
	admin.Name = name
	pwd := this.GetString("pwd")
	pwd2 := this.GetString("pwd2")
	if pwd == "" || pwd2 == "" || pwd != pwd2 {
		this.ReturnJson(map[string]string{"message":"密码格式不正确"},400)
	}
	pwdMd5 := funcs.MakeMd5(pwd+"yan")
	admin.Password = pwdMd5
	mobile := this.GetString("mobile")
	if mobile == "" {

		this.ReturnJson(map[string]string{"message":"手机号不正确"},400)
	}
	exist := orm.NewOrm().QueryTable("admin").Filter("mobile",mobile).Exist()
	if exist {
		this.ReturnJson(map[string]string{"message":"手机号已存在"},400)
	}
	admin.Mobile = mobile

	_,err := orm.NewOrm().Insert(&admin)

	if err != nil {
		this.ReturnJson(map[string]string{"message":"插入数据错误"+err.Error()},400)
	}

	this.ReturnJson(map[string]string{"message":"添加成功"},200)
}
/*
	用户列表 的 table 数据
 */
func(this *MemberController)ListData(){
	var admin []*model.Admin
	o := orm.NewOrm().QueryTable("admin")
	if startTime,endTime:=this.GetString("start"),this.GetString("end");startTime != "" && endTime != "" {
		startTime = startTime+" 00:00:00"
		endTime = endTime+" 23:59:59"
		//时间参数
		o = o.Filter("created_at__between",startTime,endTime)
	}

	if userName := this.GetString("username");userName != "" {
		//用户姓名
		o =o.Filter("name",userName)
	}
	o =o.Filter("deleted_at__isnull",true).OrderBy("-id")
	o,myPage,err3:=this.GetPage(o)
	if err3 != nil {
		this.ReturnJson(map[string]string{"message":"分页错误"},400)
	}

	_,err :=o.All(&admin)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"查询错误"},400)
	}

	var adminData []*model.AdminData


	for _,v := range admin {
		sexName := ""
		switch v.Sex {
			case 0:sexName = "未知"
			case 1:sexName = "男"
			case 2:sexName = "女"
			default:sexName = ""
		}
		adminData = append(adminData,&model.AdminData{v,sexName})
	}
	this.ReturnJson(map[string]interface{}{"code":0,"data":adminData,"count":myPage.Count},200)
}


func(this *MemberController)Del(){
	adminId := this.GetString("id")
	if adminId == "" {
		this.ReturnJson(map[string]string{"message":"请传入id"},400)
	}
	id,err := strconv.Atoi(adminId)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"请传入数字"},400)
	}
	delTime := time.Now().Format("2006-01-02 15:04:05")
	num,err := orm.NewOrm().QueryTable("admin").Filter("id",id).Update(orm.Params{
		"deleted_at":delTime,
	})

	if err != nil && num > 0 {
		this.ReturnJson(map[string]string{"message":"删除失败"},400)
	}
	this.ReturnJson(map[string]string{"message":"删除成功"},200)


}
