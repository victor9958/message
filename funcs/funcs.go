package funcs

import (
	"crypto/md5"
	"fmt"
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
func Emplode(s string,del string)[]string{
	//u_s := []rune(s)
	//u_del := []rune(del)
	var arr []string
	temp := ""
	for _,v := range s{

		if  string(v) == del{
			arr  = append(arr,temp)
			temp = ""
		}else{
			temp = temp+string(v)
		}
	}
	if temp !="" {
		arr  = append(arr,temp)
	}
	return arr
}


