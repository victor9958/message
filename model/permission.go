package model

import "time"

type Permissions struct {
	Id int             		`json:"id"`
	Title string     		`json:"title"`
	AliasTitle string 		`json:"alias_title"`
	Rule string 			`json:"rule"`
	Method string 			`json:"method"`
	Icon string 			`json:"icon"`
	Type int 				`json:"type"`
	ParentId int 			`json:"parent_id"`
	Sort int		 		`json:"sort"`
	Default int 			`json:"default"`
	CreatedAt time.Time			`json:"created_at"`
	UpdatedAt time.Time			`json:"updated_at"`
	DeletedAt time.Time			`json:"deleted_at"`
}


type PermissionsNode struct {
	Permissions
	PermissionsNodes 	[]*PermissionsNode 	`json:"children"`
}

func Walk(pNode *PermissionsNode)  {

}




