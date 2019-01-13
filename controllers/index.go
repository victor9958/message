package controllers

import (
	"github.com/astaxie/beego"
	"message/model"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	beego.Info(	string(model.MyRedis.Get("time").([]byte)))
	beego.Info("sssssssssss")

	this.TplName = "index.html"
}


