package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"message/model"
	"strconv"
	"time"
)

func IsLogin(ctx *context.Context) {
	timeStr := string(model.MyRedis.Get("time").([]byte))
	if timeStr == "" {
		ctx.Redirect(302,"/login-page")
	}
	timeInt,err  := strconv.Atoi(timeStr)
	if err!=nil {
		beego.Error(err)
		ctx.Redirect(302,"/login-page")
	}
	nowTime := time.Now().Unix()
	beego.Info(nowTime)
	if res := nowTime - int64(timeInt+3600*24);res>0 {
		ctx.Redirect(302,"/login-page")
	}
}