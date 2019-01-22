package model

import (
	//"github.com/astaxie/beego/validation"
	"time"
)

type Permissions struct {
	Id int             		`json:"id"`
	Title string     		`json:"title"`
	AliasTitle string 		`json:"alias_title"`
	Rule string 			`json:"rule"`
	Method string 			`json:"method"`
	Icon string 			`json:"icon"`
	Type int 				`json:"type"`
	ParentId int 			`json:"parent_id"`
	Level int 				`json:"level"`
	Sort int		 		`json:"sort"`
	Default int 			`json:"default"`
	CreatedAt time.Time			`json:"created_at"`
	UpdatedAt time.Time			`json:"updated_at"`
	DeletedAt time.Time			`json:"deleted_at"`
}


type PermissionsNode struct {
	Permissions
	PermissionsNodes 	*[]PermissionsNode 	`json:"children"`
}
// level  = 0
//func MakeTree(pNode []*PermissionsNode,p []*Permissions,level int) {
//	if len(p)==0 {
//		return
//	}
//	for k,v:= range p{
//		if level == v.Level{
//
//			p = append(p[:k],p[k+1])
//			tmpP := []*PermissionsNode{}
//			for _,vv := range p{
//				if v.Id == vv.ParentId {
//					tmpP = append(tmpP,&PermissionsNode{*vv,[]*PermissionsNode{}})
//				}
//			}
//			pNode = append(pNode,&PermissionsNode{*v,tmpP})
//
//		}
//	}
//
//}



func BuildData(list []*PermissionsNode) map[int]map[int]*PermissionsNode {
	var data map[int]map[int]*PermissionsNode = make(map[int]map[int]*PermissionsNode)
	for _, v := range list {
		id := v.Id
		fid := v.ParentId
		if _, ok := data[fid]; !ok {
			data[fid] = make(map[int]*PermissionsNode)
		}
		data[fid][id] = v
	}
	return data
}

func MakeTreeCore(index int, data map[int]map[int]*PermissionsNode) []*PermissionsNode {
	tmp := make([]*PermissionsNode, 0)
	//for id, item := range data[index] {
	//	if data[id] != nil {
	//		item.PermissionsNodes = MakeTreeCore(id, data)
	//	}
	//	tmp = append(tmp, item)
	//}
	return tmp
}


type List struct {
	Id       int
	Title    string
	ParentId int
	Icon     string
	Path     string
	Children *[]List
}

func MakeTree(Trees []Permissions, parentId int, nodeList *[]PermissionsNode) {
	for _, val := range Trees {

		if parentId == val.ParentId {

			temp := make([]PermissionsNode, 0)

			child := PermissionsNode{val,  &temp}

			*nodeList = append(*nodeList, child)

			MakeTree(Trees, val.Id, &temp)
		}
	}
}

