package model

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	user := beego.AppConfig.String("mysqluser")
	passwd := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	port,err := beego.AppConfig.Int("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")

	if err!=nil {
		port =3306
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		user, passwd, host, port, dbname))

	orm.RegisterModel(new(Admin))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	orm.RunSyncdb("default", false, false)

}