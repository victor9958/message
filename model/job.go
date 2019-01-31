package model

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Job struct {
	Id int             		`json:"id"`
	//Uuid string     		`json:"uuid"`
	Name string     		`json:"name"`
	RoleIds string     		`json:"role_ids"`
	CreatedAt time.Time			`json:"created_at"`
	UpdatedAt time.Time			`json:"updated_at"`
	DeletedAt time.Time			`json:"deleted_at"`

}


func GetJobById(id int)(Job,error){
	var job Job
	var err error
	err = orm.NewOrm().QueryTable("job").
		Filter("id",id).One(&job)
	return job,err
}




