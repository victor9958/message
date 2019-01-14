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


