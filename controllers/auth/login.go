package auth

import (
	"fmt"
	"managIncident/models"

	"managIncident/controllers/admin"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Login() {
	flash := beego.ReadFromRequest(&this.Controller)
	if _, ok := flash.Data["notice"]; ok {
	}

	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")

	newPass := admin.Md5Pass(password)

	fmt.Println(username, newPass)

	if this.Ctx.Input.Method() == "POST" {
		o := orm.NewOrm()
		o.Using("default")

		var users []*models.User

		qs := o.QueryTable("user")
		err := qs.Filter("mail__iexact", username).Filter("pass__iexact", newPass).One(&users)
		if err == orm.ErrNoRows {
			// No result
			flash := beego.NewFlash()
			flash.Error("Vérifie tes informations car apparemment elles ne sont pas bonnes.")
			flash.Store(&this.Controller)
		} else {
			v := this.GetSession("IncidentManager")
			if v == nil {
				this.SetSession("IncidentID", int(1))
				for _, data := range users {
					this.SetSession("uid", data.Id)
					this.SetSession("mail", data.Mail)
					this.SetSession("role", data.Role)
				}
				this.Data["num"] = 0

			} else {
				this.SetSession("IncidentID", v.(int)+1)
				this.Data["num"] = v.(int)
			}
			flash := beego.NewFlash()

			flash.Notice("Bienvenue : " + username)
			flash.Store(&this.Controller)
			if this.GetSession("role") == "admin" {
				this.Redirect("/incident-manager/admin", 302)
			} else if this.GetSession("role") == "user" {
				this.Redirect("/incident-manager", 302)
			}

		}

	}

	this.Layout = "layout.tpl"
	this.TplNames = "login.tpl"
	this.Data["title2"] = "Se connecter"
	this.Data["role"] = this.GetSession("role")
	this.Data["mail"] = this.GetSession("mail")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["navbar"] = "index/navbar.tpl"
	this.LayoutSections["footer"] = "index/footer.tpl"

}
