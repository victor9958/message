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
	var permissions []model.Permissions
	var p2 []model.PermissionsNode
	_,err := orm.NewOrm().QueryTable("permissions").Filter("type",1).OrderBy("-id").All(&permissions)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"列表查询出错"+err.Error()},400)
	}
	model.MakeTree(permissions,0,&p2)

	this.Data["data"] = p2
	this.TplName = "index.html"
}


