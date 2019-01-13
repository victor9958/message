package model

import "time"

type Admin struct {
	Id int             		`json:"id"`
	Uuid string     		`json:"uuid"`
	Name string     		`json:"name"`
	Email string   			`json:"email"`
	Mobile string   		`json:"mobile"`
	JobNo string   			`json:"job_no"`
	Password string 		`json:"password"`
	Remark string			`json:"remark"`
	RoleId int				`json:"role_id"`
	Sort int		 		`json:"sort"`
	Sex int					`json:"sex"`
	Super int				`json:"super"`
	Status int				`json:"status"`
	Position string			`json:"position"`
	CampusId int		 	`json:"campus_id"`
	CanteenId int		 	`json:"canteen_id"`
	CreatedAt time.Time			`json:"created_at"`
	UpdatedAt time.Time			`json:"updated_at"`
	DeletedAt time.Time			`json:"deleted_at"`

}

type AdminData struct {
	*Admin
	SexName string 			`json:"sex_name"`
}



