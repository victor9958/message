package main

import (
	"github.com/astaxie/beego"

	_ "message/routers"
)

func main() {
	//beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.Run()
}

