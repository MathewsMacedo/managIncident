package admin

import (
	"fmt"
	"managIncident/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

func (this *AdminController) GetUser() {

	flash := beego.ReadFromRequest(&this.Controller)
	if _, ok := flash.Data["notice"]; ok {
	}

	badgeCount(this)

	o := orm.NewOrm()
	o.Using("default")

	var users []*models.User

	qs := o.QueryTable("User")
	num, err := qs.OrderBy("-id").All(&users)

	if err != orm.ErrNoRows && num > 0 {
		this.Data["users"] = users

	} else {
		flash.Error("Aucun Utilisateur dans la Base de données")
		flash.Store(&this.Controller)
		this.Redirect("/incident-manager/admin/user", 302)
	}

	Template(this, "admin", "user", "Tous les Utilisateurs")
}

func (this *AdminController) AddUser() {
	o := orm.NewOrm()
	o.Using("default")

	var users models.User
	var register models.Register
	flash := beego.NewFlash()

	err := this.ParseForm(&users)
	if err != nil {
		beego.Error("Impossible de Parser", err)
	} else {
		valid := validation.Validation{}

		valid.Required(users.Mail, "mail")
		valid.Required(users.Role, "role")

		isValid, _ := valid.Valid(users)

		if this.Ctx.Input.Method() == "POST" {

			if !isValid {
				flash.Error("Les informations rentrées ne sont pas bonnes")
				flash.Store(&this.Controller)
				this.Redirect("/incident-manager/admin/user", 302)

				for _, err := range valid.Errors {
					beego.Error(err.Key, err.Message)
				}
			} else {
				id, _ := strconv.Atoi(this.Ctx.Request.FormValue("register_id"))
				mail := this.Ctx.Request.FormValue("mail")
				role := this.Ctx.Request.FormValue("role")
				register = models.Register{Id: id}

				users = models.User{Mail: mail, Md5Mail: Md5Pass(mail), Role: role, Register: &register}
				_, err := o.Insert(&users)
				if err == nil {
					//Cryptage email pour lien validation
					newMail := Md5Pass(mail)

					//sendMail
					err := sendMail(mail, newMail)
					if err != nil {
						flash.Error("Utilisateur enregistré mais Email Non envoyé à : " + mail)
						flash.Store(&this.Controller)

					} else {
						flash.Success("Utilisateur enregistré et Email bien envoyé à : " + mail)
						flash.Store(&this.Controller)
					}
					this.Redirect("/incident-manager/admin/user", 302)
				} else {
					flash.Error("Cet utilisateur n'a pu être rajouter")
					flash.Store(&this.Controller)
					this.Redirect("/incident-manager/admin/register", 302)
				}

			}

		}
	}

	Template(this, "admin", "user", "Validation Utilisateur")

}

func (this *AdminController) EditUser() {
	o := orm.NewOrm()
	o.Using("default")

	usersId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	users := models.User{}

	flash := beego.NewFlash()

	err := o.QueryTable("user").Filter("id", usersId).One(&users)

	if err != orm.ErrNoRows {

		err := this.ParseForm(&users)
		if err != nil {

			beego.Error("Impossible de parser. Raison: ", err)

		} else {

			valid := validation.Validation{}

			valid.Required(users.Mail, "mail")
			valid.Required(users.Role, "role")

			isValid, _ := valid.Valid(users)

			if this.Ctx.Input.Method() == "POST" {

				if !isValid {
					flash.Error("Impossible de mettre à jour l'utilisateur")
					flash.Store(&this.Controller)
					this.Redirect("/incident-manager/admin/user", 302)
				} else {
					_, err := o.Update(&users)

					if err == nil {
						flash.Notice("Utilisateur " + users.Mail + " mis à jour")
						flash.Store(&this.Controller)

						this.Redirect("/incident-manager/admin/user", 302)
					} else {
						fmt.Println("erreur")

						beego.Debug("Mise à jour Impossible dû a : ", err)
					}
				}

			}

		}
		this.Redirect("/incident-manager/admin/user", 302)

	} else {
		flash.Notice("Utilisateur %d n'existe pas", usersId)
		flash.Store(&this.Controller)
		this.Redirect("/incident-manager/", 302)
	}

}

func (this *AdminController) DeleteUser() {
	o := orm.NewOrm()
	o.Using("default")

	usersId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	users := models.User{}

	flash := beego.NewFlash()

	if exist := o.QueryTable(users.TableName()).Filter("Id", usersId).Exist(); exist {
		if num, err := o.Delete(&models.User{Id: usersId}); err == nil {
			beego.Info("Record Deleted. ", num)
			flash.Warning("L'utilisateur a bien été supprimé")
		} else {
			beego.Error("L'utilisateur n'a pu être supprimé. Raison: ", err)
		}

	} else {
		flash.Error("L'utilisateur n'existe pas %d", usersId)
	}

	flash.Store(&this.Controller)

	this.Redirect("/incident-manager/admin/user", 302)
}
