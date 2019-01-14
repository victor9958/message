package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/session"
	"message/model"
	"strconv"
	"time"
)

func IsLogin(ctx *context.Context) {
	beego.Info("isLogin")
	//string(model.MyRedis.Get("time").([]byte))
	userId := ctx.Input.CruSession.Get("user_id")
	if userId == nil {
		beego.Info("user_id==nil")
		ctx.Redirect(302,"/login-page")
		return
	}
	beego.Info(userId)
	id,ok := userId.(int)

	idStr := strconv.Itoa(id)
	if !ok {
		beego.Info("isLogin_ok")
		ctx.Redirect(302,"/login-page")
		return
	}
	timeByte := model.MyRedis.Get("time:"+idStr)
	beego.Info("time_byte")
	if timeByte == nil {
		beego.Info("time_byte_redirect")
		ctx.Redirect(302,"/login-page")
		return
	}
	timeStr := string(timeByte.([]byte))
	beego.Info("time_str")

	if timeStr == "" {
		ctx.Redirect(302,"/login-page")
		return
	}
	timeInt,err  := strconv.Atoi(timeStr)
	if err!=nil {
		beego.Error(err)
		ctx.Redirect(302,"/login-page")
		return
	}
	nowTime := time.Now().Unix()
	beego.Info(nowTime)
	//if res := nowTime - int64(timeInt+3600*24);res>0 {
	if res := nowTime - int64(timeInt+24*3600);res>0 {
		ctx.Redirect(302,"/login-page")
		return
	}
	beego.Info("isLogin_end")
}