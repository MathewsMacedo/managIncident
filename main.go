package main

import (
	"fmt"
	_ "managIncident/controllers/admin"
	models "managIncident/models"
	_ "managIncident/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver(beego.AppConfig.String("data_type"), orm.DR_MySQL)
	orm.RegisterDataBase("default", beego.AppConfig.String("data_type"), beego.AppConfig.String("data_user")+":"+beego.AppConfig.String("data_pass")+"@/"+beego.AppConfig.String("data_db")+"?charset=utf8&loc=Europe%2FParis")
	orm.RegisterModel(new(models.Incident), new(models.User), new(models.Register))
}

func main() {
	name := "default"
	force := false
	verbose := true
	if beego.SessionOn != true {
		beego.SessionOn = true
		beego.SessionName = "IncidentManager"
		beego.SessionProvider = "file"
		beego.SessionSavePath = "./tmp"
	}

	err := orm.RunSyncdb(name, force, verbose)
	orm.Debug = false
	if err != nil {
		beego.Debug(err)
	}

	beego.AdminHttpAddr = "localhost"
	beego.AdminHttpPort = 8888
	beego.Run()
	go fmt.Println("Webservice Started")

}
