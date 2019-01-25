package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"message/funcs"
	"message/model"
	"strconv"
	"strings"
	"time"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	beego.Info("index_startS")
	//获取ｊｏｂ列表
	idInter := this.GetSession("user_id")
	if idInter == nil {
		this.ReturnJson(map[string]string{"message":"列表查询出错"},400)
	}
	id,ok := idInter.(int)
	if !ok {
		this.ReturnJson(map[string]string{"message":"列表查询出错"},400)
	}

	var permissions []model.Permissions
	var p2 []model.PermissionsNode
	var admin model.Admin
	err := orm.NewOrm().QueryTable("admin").Filter("id",id).One(&admin)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"列表查询出错"+err.Error()},400)
	}
	if  admin.JobIds == ""{
		this.ReturnJson(map[string]string{"message":"列表查询出错"+err.Error()},400)
	}
	arr := funcs.Emplode(admin.JobIds,",")
	var jobs []model.Job
	num,err2 := orm.NewOrm().QueryTable("job").Filter("id__in",arr).All(&jobs)
	if err2 != nil  || num == 0 {
		this.ReturnJson(map[string]string{"message":"列表查询出错"+err.Error()},400)
	}
	permissionStr := ""
	for _,v := range jobs{
		permissionStr +=v.RoleIds
	}
	arr2 := strings.Split(permissionStr,",")
	arr3 := funcs.GetIdArr(arr2)


	_,err4 := orm.NewOrm().QueryTable("permissions").Filter("id__in",arr3).
		Filter("type",1).OrderBy("-id").All(&permissions)
	if err4 != nil {
		this.ReturnJson(map[string]string{"message":"列表查询出错"+err.Error()},400)
	}
	var urls []string
	for _,v := range permissions{
		urls = append(urls,v.Rule)
	}

	urlsJson,err3 := json.Marshal(urls)

	if err3 != nil {
		this.ReturnJson(map[string]string{"message":"列表查询出错"+err.Error()},400)
	}
	beego.Info(string(urlsJson))
	model.MyRedis.Put("urls:"+strconv.Itoa(id),string(urlsJson),1000*time.Second)

	model.MakeTree(permissions,0,&p2)

	this.Data["data"] = p2
	this.TplName = "index.html"
}


