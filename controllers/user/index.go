package user

import (
	"fmt"
	"managIncident/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego/validation"
)

//Index Controller
type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	flash := beego.ReadFromRequest(&this.Controller)
	if _, ok := flash.Data["notice"]; ok {
	}
	v := this.GetSession("IncidentManager")
	if v == nil {
		this.SetSession("IncidentID", int(1))
		this.Data["num"] = 0

	} else {
		this.SetSession("IncidentID", v.(int)+1)
		this.Data["num"] = v.(int)
	}

	r := this.Ctx.Input
	fmt.Println("Utilisateur avec comme IP : " + r.IP())

	o := orm.NewOrm()
	o.Using("default")

	var incidents []*models.Incident

	num, err := o.QueryTable("Incident").OrderBy("-id").RelatedSel().All(&incidents)

	if err != orm.ErrNoRows && num > 0 {
		this.TplNames = "index/index.tpl"
		this.Data["incidents"] = incidents

	} else {
		// No result
		flash.Error("Aucun incident dans la base de données")
		flash.Store(&this.Controller)
	}

	Template(this, "index", "index.tpl", "Liste incidents")

}

func (this *IndexController) ByMail() {
	flash := beego.NewFlash()
	o := orm.NewOrm()
	o.Using("default")

	var incidents []*models.Incident

	mail := this.GetSession("uid")

	num, err := o.QueryTable("Incident").Filter("User", mail).RelatedSel().All(&incidents)

	if err != orm.ErrNoRows && num > 0 {

		this.TplNames = "index/index.tpl"
		this.Data["incidents"] = incidents

		flash.Notice("Mes incidents déclarés")
		flash.Store(&this.Controller)

	} else {
		// No result
		flash.Error("Aucun incident ne correspond")
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
	}

	Template(this, "user", "myincident.tpl", "Liste de mes incidents déclarés")

}

func (this *IndexController) Add() {

	o := orm.NewOrm()
	o.Using("default")

	incidents := models.Incident{}
	flash := beego.NewFlash()
	// this.Data["Form"] = &incidents

	if err := this.ParseForm(&incidents); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {
		valid := validation.Validation{}

		valid.Required(incidents.Title, "title")
		valid.Required(incidents.Cat, "cat")
		valid.Required(incidents.Description, "description")
		valid.Required(incidents.DateRequest, "dateRequest")
		valid.Required(incidents.Priority, "priority")
		t := incidents.DateRequest
		date := t.Format("2006-01-02 15:04:05")
		isValid, _ := valid.Valid(incidents)

		if this.Ctx.Input.Method() == "POST" {

			if !isValid {
				this.Data["errors"] = valid.ErrorsMap

				for _, err := range valid.Errors {
					beego.Error(err.Key, err.Message)
				}

			} else {
				// _, err := o.Insert(&incidents)

				res, err := o.Raw("INSERT INTO `incident` (`cat`, `title`, `description`,`date_request`, `priority`, `user_id`) VALUES (?,?,?,?,?,?)", incidents.Cat, incidents.Title, incidents.Description, date, incidents.Priority, this.GetSession("uid")).Exec()
				if err == nil {
					num, _ := res.RowsAffected()
					fmt.Println("mysql row affected nums: ", num)
					flash.Notice("Incident " + incidents.Title + " added")
					flash.Store(&this.Controller)
					this.Redirect("/incident-manager/", 302)
				} else {
					beego.Debug("Couldn't insert new incident. Reason: ", err)
				}
			}

		}
		this.Data["dateRequest"] = time.Now()
		Template(this, "user", "declaration.tpl", "Déclarer un incident")
	}
	// defer this.DestroySession()
}

func (this *IndexController) GetOne() {

	flash := beego.NewFlash()
	o := orm.NewOrm()
	o.Using("default")

	// Get the ID page
	incidentsId := this.Ctx.Input.Param(":id")

	var incidents []*models.Incident

	err := o.QueryTable("incident").Filter("id", incidentsId).RelatedSel().One(&incidents)

	if err == orm.ErrNoRows {
		// No result
		flash.Error("Aucun incident ne correspond")
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
	} else {
		this.TplNames = "incident/incidentOne.tpl"
		for _, data := range incidents {
			this.Data["id"] = data.Id
			this.Data["cat"] = data.Cat
			this.Data["title"] = data.Title
			this.Data["description"] = data.Description
			this.Data["resolv"] = data.Resolv
			this.Data["dateRequest"] = data.DateRequest
			this.Data["dateEstimated"] = data.DateEstimated
			this.Data["dateResolution"] = data.DateResolution
			this.Data["priority"] = data.Priority
			this.Data["confirmUser"] = data.ConfirmUser
			this.Data["user"] = data.User
		}

	}

	Template(this, "incident", "incidentOne.tpl", "Incident")

}

func (this *IndexController) Edit() {
	o := orm.NewOrm()
	o.Using("default")

	incidentId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	incidents := models.Incident{}

	flash := beego.NewFlash()

	err := o.QueryTable("incident").Filter("id", incidentId).One(&incidents)

	if err != orm.ErrNoRows {

		err := this.ParseForm(&incidents)
		if err != nil {

			beego.Error("Impossible de parser. Raison: ", err)

		} else {
			timer := time.NewTimer(time.Second * 60)
			stop := timer.Stop()
			if this.Ctx.Input.Method() == "POST" {
				confirmUser := this.Ctx.Input.Request.FormValue("confirmUser")
				if confirmUser == "2" || confirmUser == "3" {
					incidents.DateResolution = time.Time{}
				}
				num, err := o.Update(&incidents)
				if err == nil {
					if num > 0 {

						if confirmUser == "1" {

							flash.Warning("Incident : " + incidents.Title + " terminé. Merci à toi.")
							flash.Store(&this.Controller)

						} else if confirmUser == "2" || confirmUser == "3" {
							fmt.Println("Stop timer", stop)
							fmt.Println("mysql row affected nums: ", num)
							flash.Warning("Incident : " + incidents.Title + " de nouveau renvoyé à l'administrateur")
							flash.Store(&this.Controller)
						}

					} else {
						flash.Error("Rien a été modifié pour l'incident : " + incidents.Title)
						flash.Store(&this.Controller)
					}

					this.Redirect("/incident-manager/", 302)

				} else {
					fmt.Println("erreur")

					beego.Debug("Mise à jour Impossible dû a : ", err)
				}

			}

		}
		this.Redirect("/incident-manager/", 302)

	} else {
		flash.Error("Incident %d n'existe pas", incidentId)
		flash.Store(&this.Controller)
		this.Redirect("/incident-manager/", 302)
	}

}

func Template(this *IndexController, dossier string, tpl string, titre string) {
	this.Data["mail"] = this.GetSession("mail")
	this.Data["role"] = this.GetSession("role")
	this.Data["uid"] = this.GetSession("uid")
	this.Data["dateNow"] = time.Now()
	this.Layout = "layout.tpl"
	this.TplNames = dossier + "/" + tpl
	this.Data["title2"] = titre
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["navbar"] = "index/navbar.tpl"
	this.LayoutSections["footer"] = "index/footer.tpl"
}
