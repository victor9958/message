package funcs

import (
	"crypto/md5"
	"fmt"
)

func MakeMd5(str string)string{
	has := md5.Sum([]byte(str))
	return fmt.Sprintf("%x",has)
}
