package controllers

import "message/funcs"

type CeshiController struct {
	BaseController
}

func(this *CeshiController)Ceshi(){
	pwd := funcs.MakeMd5("123456")
	this.ReturnJson(pwd,200)
}
