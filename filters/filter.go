package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session"
	"message/funcs"
	"message/model"
	"strconv"
	"strings"
	"time"
)

func IsLogin(ctx *context.Context) {
	beego.Info("isLogin")
	//string(model.MyRedis.Get("time").([]byte))
	userId := ctx.Input.CruSession.Get("user_id")
	if userId == nil {
		ctx.Redirect(302,"/login-page")
		return
	}
	id,ok := userId.(int)

	idStr := strconv.Itoa(id)
	if !ok {
		ctx.Redirect(302,"/login-page")
		return
	}
	timeByte := model.MyRedis.Get("time:"+idStr)
	if timeByte == nil {
		ctx.Redirect(302,"/login-page")
		return
	}
	timeStr := string(timeByte.([]byte))

	if timeStr == "" {
		ctx.Redirect(302,"/login-page")
		return
	}
	timeInt,err  := strconv.Atoi(timeStr)
	if err!=nil {
		ctx.Redirect(302,"/login-page")
		return
	}
	nowTime := time.Now().Unix()
	//if res := nowTime - int64(timeInt+3600*24);res>0 {
	if res := nowTime - int64(timeInt+24*3600);res>0 {
		ctx.Redirect(302,"/login-page")
		return
	}
}

func Auth(ctx *context.Context){
	beego.Info("auth_start")
	userId := ctx.Input.CruSession.Get("user_id")
	if userId == nil {
		ctx.Redirect(302,"/login-page")
		return
	}
	id,ok := userId.(int)
	if !ok {
		ctx.Redirect(302,"/login-page")
		return
	}
	var admin model.Admin
	err := orm.NewOrm().QueryTable("admin").Filter("id",id).One(&admin)
	if err != nil {
		ctx.Redirect(302,"/login-page")
		return
	}
	if  admin.JobIds == ""{
		ctx.Redirect(302,"/login-page")
		return
	}
	arr := funcs.Emplode(admin.JobIds,",")
	var jobs *[]model.Job
	num,err2 := orm.NewOrm().QueryTable("job").Filter("id__in",arr).All(&jobs)
	if err2 != nil  || num == 0 {
		ctx.Redirect(302,"/login-page")
		return
	}
	permissionStr := ""
	for _,v := range *jobs{
		permissionStr +=v.RoleIds
	}
	arr2 := strings.Split(permissionStr,",")
	arr3 := funcs.GetIdArr(arr2)

	var urls orm.ParamsList
	num,err2 := orm.NewOrm().QueryTable("permissions").Filter("id__in",arr3).ValuesFlat(&urls,"rule")

	model.MyRedis.Put("time:"+idStr,time.Now().Unix(),1000*time.Second)



}


