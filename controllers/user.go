package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"message/model"
	"strconv"
	"time"
)

type Usercontroller struct {
	BaseController
}

/*
	用户列表 ui框架
 */
func(this *Usercontroller)List(){
	this.TplName="user-list.html"
}

func(this *Usercontroller)ListData(){
	var users []*model.User
	o := orm.NewOrm().QueryTable("user")
	if startTime,endTime:=this.GetString("start"),this.GetString("end");startTime != "" && endTime != "" {
		startTime = startTime+" 00:00:00"
		endTime = endTime+" 23:59:59"
		//时间参数
		o = o.Filter("created_at__between",startTime,endTime)
	}

	if name := this.GetString("name");name != "" {
		//用户姓名
		o =o.Filter("name",name)
	}
	o =o.Filter("deleted_at__isnull",true).OrderBy("-id")
	o,myPage,err3:=this.GetPage(o)
	if err3 != nil {
		this.ReturnJson(map[string]string{"message":"分页错误"},400)
	}

	_,err :=o.All(&users)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"查询错误"},400)
	}
	var userData []*model.UserDetail

	for _,v := range users{
		sexName := ""
		switch v.Sex {
		case 0:sexName = "未知"
		case 1:sexName = "男"
		case 2:sexName = "女"
		default:sexName = ""
		}
		userData = append(userData,&model.UserDetail{v,sexName})
	}


	this.ReturnJson(map[string]interface{}{"code":0,"data":userData,"count":myPage.Count},200)
}


func(this *Usercontroller)AddPage(){
	this.TplName = "user-add.html"
}


func(this *Usercontroller)Add(){
	var user model.User
	user.Name = this.GetString("name")
	if user.Name  == "" {
		this.ReturnJson(map[string]string{"message":"姓名必填"},400)
	}
	user.Email = this.GetString("email")
	user.Mobile = this.GetString("mobile")
	if user.Mobile=="" {
		this.ReturnJson(map[string]string{"message":"手机号必填"},400)
	}
	sort := this.GetString("sort")
	if sort =="" {
		user.Sort = 0
	}else{
		var err error
		user.Sort,err = strconv.Atoi(sort)
		if err != nil {
			user.Sort = 0
		}
	}
	user.Remark = this.GetString("remark")
	user.Address = this.GetString("address")
	sex := this.GetString("sex")
	if sex =="" {
		user.Sex = 0
	}else{
		var err error
		user.Sex,err = strconv.Atoi(sex)
		if err != nil {
			user.Sex = 0
		}
	}
	time := time.Now()//.Format(consts.YMDHIS)
	beego.Info(time)
	user.CreatedAt = time
	user.UpdatedAt = time

	_,err := orm.NewOrm().Insert(&user)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"添加失败"+err.Error()},400)
	}
	this.ReturnJson(map[string]string{"message":"添加成功"},200)
}

func(this *Usercontroller)EditPage(){
	id,err := this.GetStringChangeInt("id")
	if err != nil {
		this.ReturnJson(map[string]string{"message":err.Error()},400)
	}
	user,err2 := model.GetUsserById(id)
	if err2 != nil {
		this.ReturnJson(map[string]string{"message":err2.Error()},400)
	}
	this.Data["id"] = id
	this.Data["user"] = user
	this.TplName = "user-edit.html"

}



func(this *Usercontroller)Edit(){
	id,err := this.GetStringChangeInt("id")
	if err != nil {
		this.ReturnJson(map[string]string{"message":err.Error()},400)
	}

	user,err5 := model.GetUsserById(id)
	if err5 != nil {
		this.ReturnJson(map[string]string{"message":"数据不存在"},400)
	}

	user.Name = this.GetString("name")
	if user.Name  == "" {
		this.ReturnJson(map[string]string{"message":"姓名必填"},400)
	}
	user.Email = this.GetString("email")
	user.Mobile = this.GetString("mobile")
	if user.Mobile=="" {
		this.ReturnJson(map[string]string{"message":"手机号必填"},400)
	}
	sort := this.GetString("sort")
	if sort =="" {
		user.Sort = 0
	}else{
		var err error
		user.Sort,err = strconv.Atoi(sort)
		if err != nil {
			user.Sort = 0
		}
	}
	user.Remark = this.GetString("remark")
	user.Address = this.GetString("address")
	sex := this.GetString("sex")
	if sex =="" {
		user.Sex = 0
	}else{
		var err error
		user.Sex,err = strconv.Atoi(sex)
		if err != nil {
			user.Sex = 0
		}
	}
	time := time.Now()//.Format(consts.YMDHIS)
	user.UpdatedAt = time

	_,err6 := orm.NewOrm().Update(&user)
	if err6 != nil {
		this.ReturnJson(map[string]string{"message":"修改失败"+err6.Error()},400)
	}
	this.ReturnJson(map[string]string{"message":"修改成功"},200)
}
