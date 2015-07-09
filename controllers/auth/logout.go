package auth

import "github.com/astaxie/beego"

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Logout() {

	this.DestroySession()
	flash := beego.NewFlash()

	flash.Notice("A bientôt")
	flash.Store(&this.Controller)
	this.Redirect("/incident-manager/", 302)

}
