package admin

import (
	"managIncident/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func (this *AdminController) RegisterDemand() {

	flash := beego.ReadFromRequest(&this.Controller)
	if _, ok := flash.Data["notice"]; ok {
	}

	badgeCount(this)

	o := orm.NewOrm()
	o.Using("default")

	var demand []*models.Register

	num, err := o.QueryTable("Register").OrderBy("-id").RelatedSel().All(&demand)

	if err != orm.ErrNoRows && num > 0 {
		this.TplNames = "admin/register.tpl"
		this.Data["demand"] = demand

	} else {
		// No result
		flash.Error("Aucune demande dans la Base de données")
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
	}

	Template(this, "admin", "register", "Demande de connexion")
}

func (this *AdminController) DeleteDemand() {
	o := orm.NewOrm()
	o.Using("default")

	demandId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	register := models.Register{}

	flash := beego.NewFlash()

	if exist := o.QueryTable(register.TableName()).Filter("Id", demandId).Exist(); exist {
		if num, err := o.Delete(&models.Register{Id: demandId}); err == nil {
			beego.Info("Record Deleted. ", num)
			flash.Notice("La demande a bien été supprimé")
		} else {
			beego.Error("La demande n'a pu être supprimé. Raison: ", err)
		}

	} else {
		flash.Notice("La demande n'existe pas %d", demandId)
	}

	flash.Store(&this.Controller)

	this.Redirect("/incident-manager/admin/register", 302)
}
