package admin

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
type AdminController struct {
	beego.Controller
}

func (this *AdminController) GetIncident() {

	if this.GetSession("role") != "admin" {
		this.Redirect("/incident-manager", 302)
	}

	flash := beego.ReadFromRequest(&this.Controller)
	if _, ok := flash.Data["notice"]; ok {
	}

	badgeCount(this)

	r := this.Ctx.Input
	fmt.Println("Administrateur avec comme IP : " + r.IP())

	o := orm.NewOrm()
	o.Using("default")

	var incidents []*models.Incident

	num, err := o.QueryTable("Incident").OrderBy("-id").RelatedSel().All(&incidents)

	if err != orm.ErrNoRows && num > 0 {
		this.TplNames = "index/index.tpl"
		this.Data["incidents"] = incidents
	} else {
		// No result
		flash.Error("Aucun Incident dans la Base de données")
		flash.Store(&this.Controller)
	}

	Template(this, "index", "index", "Liste incidents")

}

func (this *AdminController) ByMailIncident() {
	flash := beego.NewFlash()
	o := orm.NewOrm()
	o.Using("default")

	var incidents []*models.Incident

	mail := this.GetSession("uid")

	num, err := o.QueryTable("Incident").Filter("User", mail).RelatedSel().All(&incidents)

	if err != orm.ErrNoRows && num > 0 {

		this.TplNames = "index/index.tpl"
		this.Data["incidents"] = incidents

		flash := beego.NewFlash()

		flash.Notice("Mes incidents déclarés")
		flash.Store(&this.Controller)

	} else {
		// No result
		flash.Error("Aucun Incident dans la Base de données")
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
	}

	Template(this, "user", "myincident", "Liste de mes incidents déclarés")

}

func (this *AdminController) AddIncident() {

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
					this.Redirect("/incident-manager/admin", 302)
				} else {
					beego.Debug("Couldn't insert new incident. Reason: ", err)
				}
			}

		}
		this.Data["dateRequest"] = time.Now()
		Template(this, "user", "declaration", "Déclarer un incident")
	}
	// defer this.DestroySession()
}

func (this *AdminController) GetOneIncident() {
	flash := beego.NewFlash()
	o := orm.NewOrm()
	o.Using("default")

	// Get the ID page
	incidentsId := this.Ctx.Input.Param(":id")

	var incidents []*models.Incident

	err := o.QueryTable("incident").Filter("id", incidentsId).RelatedSel().One(&incidents)

	if err == orm.ErrNoRows {
		// No result
		flash.Error("Aucun Incident ne correspond")
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

	Template(this, "incident", "incidentOne", "Incident")

}

func (this *AdminController) EditIncident() {
	o := orm.NewOrm()
	o.Using("default")

	IncidentID, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	incidents := models.Incident{}

	flash := beego.NewFlash()

	err := o.QueryTable("incident").Filter("id", IncidentID).RelatedSel().One(&incidents)

	if err != orm.ErrNoRows {

		err := this.ParseForm(&incidents)
		if err != nil {

			beego.Error("Impossible de parser. Raison: ", err)

		} else {

			if this.Ctx.Input.Method() == "POST" {
				// res, err := o.Raw("UPDATE `incident` SET `cat`=?, `description`=?, `resolv`=?, `date_estimated`=?, `date_resolution`=?, `priority`=? WHERE ?", incidents.Cat, incidents.Description, incidents.Resolv, dateE, dateR, incidents.Priority, IncidentManager).Exec()
				num, err := o.Update(&incidents)
				if err == nil {
					if num > 0 {
						fmt.Println("mysql row affected nums: ", num)
						flash.Notice("Incident : " + incidents.Title + " mis à jour ;-)")
						flash.Store(&this.Controller)
						SendMailUpdate(incidents.User.Mail, IncidentID, incidents.Title)
					} else {
						flash.Error("Rien a été modifié pour l'incident : " + incidents.Title)
						flash.Store(&this.Controller)
					}

					this.Redirect("/incident-manager/admin", 302)
				} else {
					fmt.Println("erreur")

					beego.Debug("Mise à jour impossible dû a : ", err)
				}

			}

		}
		this.Redirect("/incident-manager/admin", 302)

	} else {
		flash.Error("Incident %d n'existe pas", IncidentID)
		flash.Store(&this.Controller)
		this.Redirect("/incident-manager/", 302)
	}

}

func (this *AdminController) DeleteIncident() {
	o := orm.NewOrm()
	o.Using("default")

	incidentId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	incidents := models.Incident{}

	flash := beego.NewFlash()

	if exist := o.QueryTable(incidents.TableName()).Filter("Id", incidentId).Exist(); exist {
		if num, err := o.Delete(&models.Incident{Id: incidentId}); err == nil {
			beego.Info("Record Deleted. ", num)
			flash.Notice("L'incident a bien été supprimé")
		} else {
			beego.Error("L'incident n'a pu être supprimé. Raison: ", err)
		}

	} else {
		flash.Notice("L'incident n'existe pas %d", incidentId)
	}

	flash.Store(&this.Controller)

	this.Redirect("/incident-manager/admin/", 302)
}

func Template(this *AdminController, dossier string, tpl string, titre string) {
	this.Data["mail"] = this.GetSession("mail")
	this.Data["role"] = this.GetSession("role")
	this.Data["uid"] = this.GetSession("uid")
	this.Data["badgeDemand"] = this.GetSession("unbrdemand")
	this.Data["badgeUser"] = this.GetSession("unbrusertotal")
	this.Data["badgeIncident"] = this.GetSession("unbrincident")
	this.Data["admin"] = "admin/"
	this.Data["dateNow"] = time.Now().Local()
	this.Layout = "layout.tpl"
	this.TplNames = dossier + "/" + tpl + ".tpl"
	this.Data["title2"] = titre
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["navbar"] = "index/navbar.tpl"
	this.LayoutSections["footer"] = "index/footer.tpl"
}

func badgeCount(this *AdminController) {
	o := orm.NewOrm()
	o.Using("default")
	//Compteur Badge
	numUser, err := o.QueryTable("user").Count()
	numDemand, err := o.QueryTable("register").Count()
	numIncident, err := o.QueryTable("incident").Count()

	if err == nil {
		this.SetSession("unbrusertotal", numUser)
		this.SetSession("unbrdemand", numDemand)
		this.SetSession("unbrincident", numIncident)

	} else {
		fmt.Println(err)
	}
	//fin compteur
}
