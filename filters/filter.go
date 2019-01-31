package filters

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/astaxie/beego/session"
	"message/model"
	"strconv"
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
	path := ctx.Input.URI()
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



	urlsByte := model.MyRedis.Get("urls:"+strconv.Itoa(id))

	if urlsByte == nil {
		ctx.Redirect(302,"/login-page")
		return
	}

	urlsJson := string(urlsByte.([]byte))


	if urlsJson == "" {
		ctx.Redirect(302,"/login-page")
		return
	}

	//pathOk := strings.Contains(urlsJson,path)
	//beego.Info(pathOk)
	//if !pathOk {
	//	ctx.Redirect(302,"/login-page")
	//	return
	//}
	//有安全隐患

	//
	var urls []string
	err := json.Unmarshal([]byte(urlsJson),&urls)
	beego.Info(urls)
	if err!=nil {
		ctx.Redirect(302,"/login-page")
		return
	}

	//用数组匹配需要进行遍历
	for _,v := range urls {
		if v == path{
			return
		}
	}
	ctx.Redirect(302,"/login-page")
	return


}


