package model

import "time"

type Admin struct {
	Id int             		`json:"id"`
	Uuid string     		`json:"uuid"`
	Name string     		`json:"name"`
	JobIds string     		`json:"job_ids"`
	Email string   			`json:"email"`
	Mobile string   		`json:"mobile"`
	Password string 		`json:"password"`
	Remark string			`json:"remark"`
	Sort int		 		`json:"sort"`
	Sex int					`json:"sex"`
	Super int				`json:"super"`
	Status int				`json:"status"`
	CreatedAt time.Time			`json:"created_at"`
	UpdatedAt time.Time			`json:"updated_at"`
	DeletedAt time.Time			`json:"deleted_at"`

}

type AdminData struct {
	*Admin
	SexName string 			`json:"sex_name"`
}



