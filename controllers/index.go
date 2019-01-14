package controllers

import (
	"github.com/astaxie/beego/orm"
	"message/model"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	//获取ｊｏｂ列表
	var permissions []*model.Permissions
	_,err := orm.NewOrm().QueryTable("permissions").Filter("type",1).All(&permissions)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"列表查询出错"+err.Error()},400)
	}
	this.Data["data"] = permissions
	this.TplName = "index.html"
}



func Walk(){

}


