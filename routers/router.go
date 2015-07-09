package routers

import (
	"managIncident/controllers"
	"managIncident/controllers/admin"
	"managIncident/controllers/auth"
	"managIncident/controllers/user"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	//Filter
	beego.InsertFilter("/", beego.BeforeRouter, FilterHome)
	beego.InsertFilter("/*", beego.BeforeRouter, FilterHome)

	beego.InsertFilter("/incident-manager/admin/*", beego.BeforeRouter, FilterAdmin)
	beego.InsertFilter("/incident-manager/admin/", beego.BeforeRouter, FilterAdmin)

	beego.InsertFilter("/incident-manager/user/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/incident-manager/user/", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/incident-manager/", beego.BeforeRouter, FilterUser)

	beego.InsertFilter("/incident-manager/user/*", beego.BeforeRouter, FilterLogin)
	beego.InsertFilter("/incident-manager/admin/*", beego.BeforeRouter, FilterLogin)

	//User Router
	beego.Router("/incident-manager/", &user.IndexController{}, "post,get:Get")
	beego.Router("/incident-manager/login", &auth.LoginController{}, "post,get:Login")
	beego.Router("/incident-manager/register", &auth.RegisterController{}, "post,get:Register")
	beego.Router("/mail/confirmation/:mail([-0-9a-z]+)", &auth.RegisterController{}, "post,get:Password")
	beego.Router("/incident-manager/logout", &auth.LogoutController{}, "post,get:Logout")
	beego.Router("/incident-manager/user/declaration", &user.IndexController{}, "get,post:Add")
	beego.Router("/incident-manager/user/myincident", &user.IndexController{}, "get,post:ByMail")
	beego.Router("/incident-manager/user/incident/:id([0-9]+)", &user.IndexController{}, "get:GetOne")
	beego.Router("/incident-manager/user/incident/update/:id([0-9]+)", &user.IndexController{}, "post,get:Edit")

	// admin Router
	beego.Router("/incident-manager/admin", &admin.AdminController{}, "post,get:GetIncident")
	beego.Router("/incident-manager/admin/user", &admin.AdminController{}, "post,get:GetUser")
	beego.Router("/incident-manager/admin/user/add", &admin.AdminController{}, "post,get:AddUser")
	beego.Router("/incident-manager/admin/user/update/:id([0-9]+)", &admin.AdminController{}, "post,get:EditUser")
	beego.Router("/incident-manager/admin/user/delete/:id([0-9]+)", &admin.AdminController{}, "post,get:DeleteUser")
	beego.Router("/incident-manager/admin/declaration", &admin.AdminController{}, "post,get:AddIncident")
	beego.Router("/incident-manager/admin/myincident", &admin.AdminController{}, "post,get:ByMailIncident")
	beego.Router("/incident-manager/admin/incident/:id([0-9]+)", &admin.AdminController{}, "get:GetOneIncident")
	beego.Router("/incident-manager/admin/incident/update/:id([0-9]+)", &admin.AdminController{}, "post,get:EditIncident")
	beego.Router("/incident-manager/admin/incident/delete/:id([0-9]+)", &admin.AdminController{}, "post,get:DeleteIncident")
	beego.Router("/incident-manager/admin/register", &admin.AdminController{}, "post,get:RegisterDemand")
	beego.Router("/incident-manager/admin/register/delete/:id([0-9]+)", &admin.AdminController{}, "post,get:DeleteDemand")
	// beego.Router("/admin/incident/:id([0-9]+)", &admin.AdminController{}, "get:GetOne")

	//Static Router
	beego.SetStaticPath("/static", "static")

	//Erreurs
	beego.ErrorController(&controllers.ErrorController{})
}

var FilterHome = func(ctx *context.Context) {
	if ctx.Input.Uri() == "/" && ctx.Input.Url() != "/(incident-manager)" {
		ctx.Redirect(302, "/incident-manager/")
	}
}

var FilterAdmin = func(ctx *context.Context) {
	if ctx.Input.Session("role") == "user" {
		ctx.Redirect(302, "/incident-manager/")
	}
}

var FilterUser = func(ctx *context.Context) {
	if ctx.Input.Session("role") == "admin" {
		ctx.Redirect(302, "/incident-manager/admin/")
	}
}

var FilterLogin = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	if !ok && ctx.Input.Uri() != "/incident-manager/login" && ctx.Input.Uri() != "/incident-manager/register" {
		ctx.Redirect(302, "/incident-manager/login")
	}

}
