package controllers

import (

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"


	"strconv"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller
	isLogin bool
}
type MyPage struct {
	Count int64	//总条数
	CountPage int	//总页数
	Limit int //每页几条
	NowPage int //当前页
}

func init(){
}



//自己的重定向
func(this *BaseController)MyRedirect(url string){
	this.Redirect(url,302)
	this.StopRun()
}
//获得ip
func(this *BaseController)GetClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr,":")
	return s[0]
}
//自己的return json数据
func(this *BaseController)ReturnJson(data interface{},status int){
	this.Ctx.Output.Status = status
	this.Ctx.Output.JSON(data,true,false)
}

func(this *BaseController)Wel(){
	this.TplName = "welcome.html"
}


//分页
func(this *BaseController)GetPage(o orm.QuerySeter)(orm.QuerySeter,*MyPage,error){


	var myPage MyPage
	myPage.Limit = 10
	myPage.NowPage = 1

	var err3  error
	myPage.Count,err3 = o.Count()
	if err3!= nil {
		return nil,&myPage, err3
	}

	//总页数
	myPage.CountPage = int(myPage.Count)/myPage.Limit
	if m := int(myPage.Count)%myPage.Limit;m>0 {
		myPage.CountPage++
	}

	if limitStr := this.GetString("limit");limitStr !="" {
		var err2 error
		myPage.Limit,err2 = strconv.Atoi(limitStr)
		if err2 != nil {
			return nil,&myPage,err2
		}
	}
	if pageStr := this.GetString("page") ;pageStr != ""{
		var err error
		myPage.NowPage,err = strconv.Atoi(pageStr)
		if err != nil {
			return nil,&myPage ,err
		}
	}

	return o.Limit(myPage.Limit,(myPage.NowPage-1)*myPage.Limit),&myPage,nil
}


/*

 */
func(this *BaseController)Ceshi(){
	//生成ｔｏｋｅｎ
 	var key []byte = []byte("hello world! this is secret!")
	claims := &jwt.StandardClaims{
		NotBefore:int64(time.Now().Unix()),
		ExpiresAt:int64(time.Now().Unix() +1000),
		Issuer:"Bitch",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	ss,err := token.SignedString(key)
	if err != nil {
		this.ReturnJson(err.Error(),400)
	}

	//验证ｔｏｋｅｎ
	b1,err2 := jwt.Parse(ss, func(ss *jwt.Token) (interface{}, error) {
		return  key,nil
	})
	if err2 != nil{
		this.ReturnJson(err2.Error(),400)
	}
	beego.Info(b1)
	this.ReturnJson(ss,200)
}
