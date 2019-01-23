package funcs

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

//生成ｍｄ５
func MakeMd5(str string)string{
	has := md5.Sum([]byte(str))
	return fmt.Sprintf("%x",has)
}

//断言
//func SwitchType(interf interface{}){
//	switch v:=e.(type) {
//	case int:
//		var s int
//		s = v
//	}
//}

//字符串转数组 只能切割单个字符 或者 单个汉字 range 循环 自动转成utf-8
func Emplode(s string,del string)[]int{
	//u_s := []rune(s)
	//u_del := []rune(del)
	var arr []int
	temp := ""
	for _,v := range s{

		if  string(v) == del{
			tempInt,err := strconv.Atoi(temp)
			if err == nil {
				arr  = append(arr,tempInt)
			}
			temp = ""
		}else{
			temp = temp+string(v)
		}
	}
	if temp !="" {
		tempInt,err := strconv.Atoi(temp)
		if err == nil {
			arr  = append(arr,tempInt)
		}
	}
	return arr
}

func GetIdArr(arr []string)[]int{
	var ids []int
	for _,v := range arr{
		if v=="" || v==" "{
			continue
		}else {
			id ,err := strconv.Atoi(v)
			if err == nil {
				ids = append(ids,id)
			}
		}
	}
	return ids
}



