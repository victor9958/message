package model

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id int             		`json:"id"`
	Name string     		`json:"name"`
	Email string   			`json:"email"`
	Mobile string   		`json:"mobile"`
	Remark string			`json:"remark"`
	Sort int		 		`json:"sort"`
	Sex int					`json:"sex"`
	Address string 			`json:"address"`
	CreatedAt time.Time			`json:"created_at"`
	UpdatedAt time.Time			`json:"updated_at"`
	DeletedAt time.Time			`json:"deleted_at"`
}

type UserDetail struct {
	*User
	SexName string             		`json:"sex_name"`
}

func GetUsserById(id int)(*User,error){
	var err error
	var user User
	err = orm.NewOrm().QueryTable("user").Filter("id",id).Filter("deleted_at__isnull",true).One(&user)
	return &user,err
}



