package controllers

import (
	"github.com/astaxie/beego/orm"
	"message/model"
)

type JobController struct {
	BaseController
}

func(this *JobController)List(){
	this.TplName = "job-list.html"
}

/*
	用户列表 的 table 数据
 */
func(this *JobController)ListData(){
	var job []*model.Job
	o := orm.NewOrm().QueryTable("job")
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

	_,err :=o.All(&job)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"查询错误"},400)
	}


	this.ReturnJson(map[string]interface{}{"code":0,"data":job,"count":myPage.Count},200)
}


