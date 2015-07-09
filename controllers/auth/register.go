package auth

import (
	"fmt"
	"managIncident/controllers/admin"
	"managIncident/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Register() {

	o := orm.NewOrm()
	o.Using("default")

	register := models.Register{}
	flash := beego.NewFlash()
	// this.Data["Form"] = &register

	if err := this.ParseForm(&register); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {
		valid := validation.Validation{}

		valid.Required(register.Mail, "mail")

		isValid, _ := valid.Valid(register)

		if this.Ctx.Input.Method() == "POST" {

			if !isValid {
				this.Data["errors"] = valid.ErrorsMap

				for _, err := range valid.Errors {
					beego.Error(err.Key, err.Message)
				}

			} else {
				r := this.Ctx.Input
				register.IP = r.IP()
				fmt.Println(register.IP)

				_, err := o.Insert(&register)

				// res, err := o.Raw("INSERT INTO `incident` (`cat`, `title`, `description`,`date_request`, `priority`, `user_id`) VALUES (?,?,?,?,?,?)", register.Cat, register.Title, register.Description, date, register.Priority, this.GetSession("uid")).Exec()
				if err == nil {
					err := admin.SendMailAdmin()
					if err != nil {
						fmt.Println(err)
					}
					flash.Success(register.Mail + " : votre demande a bien été envoyé ")
					flash.Store(&this.Controller)
					this.Redirect("/", 302)
				} else {
					flash.Warning("Attention car cette adresse mail : " + register.Mail + " est déjà dans les demandes. ")
					flash.Store(&this.Controller)
					beego.Debug("Couldn't insert in tableName Register. Reason: ", err)

				}
			}

		}

	}
	Template(this, "user", "register", "Demande de Connexion / Nouveau mot de passe")
}

func (this *RegisterController) Password() {

	flash := beego.NewFlash()
	o := orm.NewOrm()

	v := this.GetSession("uid")
	if v != nil {
		flash.Error("Une session existe déjà sur cette Ordinateur. Déconnectes toi afin d'éviter tout problème")
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
	}

	o.Using("default")
	mail := this.Ctx.Input.Param(":mail")

	user := models.User{Md5Mail: mail}
	err := o.Read(&user, "Md5Mail")
	this.Data["mail"] = user.Mail
	this.Data["md5Mail"] = mail
	// Three return values：Is Created，Object Id，Error
	if err == nil && user.Pass == "" {
		if this.Ctx.Input.Method() == "POST" {
			this.Ctx.Request.ParseForm()
			password := this.Ctx.Request.Form.Get("password")
			repassword := this.Ctx.Request.Form.Get("repassword")

			if repassword == password {

				newPass := admin.Md5Pass(password)

				user := models.User{Id: user.Id, Mail: user.Mail, Role: user.Role, Pass: newPass, Created: time.Now()}

				if _, err := o.Update(&user); err == nil {
					flash.Success("Bienvenue " + user.Mail)
					flash.Store(&this.Controller)
					v := this.GetSession("IncidentManager")
					if v == nil {
						this.SetSession("IncidentID", int(1))
						this.SetSession("uid", user.Id)
						this.SetSession("mail", user.Mail)
						this.SetSession("role", user.Role)
						this.Data["num"] = 0

					} else {
						this.SetSession("IncidentID", v.(int)+1)
						this.Data["num"] = v.(int)
					}
					this.Redirect("/incident-manager/", 302)
				} else {
					fmt.Println("update", err)
				}
			}

		}

	} else {
		flash.Error("Dommage mais tu ne peux accéder à cette page")
		flash.Store(&this.Controller)
		this.Redirect("/", 302)

	}
	Template(this, "user", "password", "Enregistre ton mot de passe")

}

func Template(this *RegisterController, dossier string, tpl string, titre string) {
	this.Data["dateRequest"] = time.Now()
	this.Layout = "layout.tpl"
	this.TplNames = dossier + "/" + tpl + ".tpl"
	this.Data["title2"] = titre
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["navbar"] = "index/navbar.tpl"
	this.LayoutSections["footer"] = "index/footer.tpl"
}
