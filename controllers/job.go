package controllers

import (
	"github.com/astaxie/beego/orm"
	"message/model"
	"strconv"
	"time"
)

type JobController struct {
	BaseController
}

func(this *JobController)List(){
	this.TplName = "job-list.html"
}


func(this *JobController)AddPage(){
	this.TplName = "job-add.html"
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


func(this *JobController)Del(){
	jobId := this.GetString("id")
	if jobId == "" {
		this.ReturnJson(map[string]string{"message":"请传入id"},400)
	}
	id,err := strconv.Atoi(jobId)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"请传入数字"},400)
	}
	delTime := time.Now().Format("2006-01-02 15:04:05")
	num,err := orm.NewOrm().QueryTable("job").Filter("id",id).Update(orm.Params{
		"deleted_at":delTime,
	})

	if err != nil && num > 0 {
		this.ReturnJson(map[string]string{"message":"删除失败"},400)
	}
	this.ReturnJson(map[string]string{"message":"删除成功"},200)

}


func(this *JobController)Add(){
	var admin model.Job
	name := this.GetString("name")
	if name == "" {
		this.ReturnJson(map[string]string{"message":"名称必填"},400)
	}
	admin.Name = name
	exist := orm.NewOrm().QueryTable("job").Filter("name",name).Exist()
	if exist {
		this.ReturnJson(map[string]string{"message":name+"已存在"},400)
	}
	_,err := orm.NewOrm().Insert(&admin)

	if err != nil {
		this.ReturnJson(map[string]string{"message":"插入数据错误"+err.Error()},400)
	}

	this.ReturnJson(map[string]string{"message":"添加成功"},200)
}



