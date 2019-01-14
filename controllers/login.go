package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"message/funcs"
	"message/model"
	"strconv"
	"time"
)

type LoginController struct {
	beego.Controller
}


func(this *LoginController)Login(){
	pwd := this.GetString("pwd")
	mobile := this.GetString("mobile")

	if mobile == "" {
		this.ReturnJson(map[string]string{"message":"手机号不存在"},400)
	}
	if pwd == "" {
		this.ReturnJson(map[string]string{"message":"密码不存在"},400)
	}

	//md5加密 pwd
	//beego.Info(pwd)
	//has := md5.Sum([]byte(pwd+"yan"))
	//pwdMB := fmt.Sprintf("%x",has)


	pwdMB := funcs.MakeMd5(pwd+"yan")
	//beego.Info(pwdMB)
	//查询user表
	var admin model.Admin
	err := orm.NewOrm().QueryTable("admin").Filter("mobile",mobile).One(&admin)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"查无此用户"},400)
	}
	if admin.Password == pwdMB {
		idStr := strconv.Itoa(admin.Id)
		beego.Info("time:"+idStr)
		this.SetSession("user_mobile",mobile)
		this.SetSession("user_name",admin.Name)
		this.SetSession("user_id",admin.Id)
		this.SetSession("login_time",	time.Now().Unix())

		model.MyRedis.Put("time:"+idStr,time.Now().Unix(),1000*time.Second)
		beego.Info(time.Now().Unix())
		this.ReturnJson(map[string]string{"message":"登录成功"},200)
	}
	this.ReturnJson(map[string]string{"message":"密码错误"},400)
}

func(this *LoginController)ReturnJson(data interface{},status int){
	this.Ctx.Output.Status = status
	this.Ctx.Output.JSON(data,true,false)
}

func(this *LoginController)LoginPage(){

	this.TplName = "login.html"
}
