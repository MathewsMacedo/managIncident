package models

import "time"

//Structure de données pour la table incident dans mySql
type Incident struct {
	Id             int       `form:"-"`
	Cat            string    `form:"cat"`
	Title          string    `form:"title" valid:"MaxSize(100)"`
	Description    string    `orm:"null;type(text)" form:"description,textarea" valid:"MaxSize(400)"`
	Resolv         string    `orm:"null;type(text)" form:"resolution,textarea" valid:"MaxSize(400)"`
	DateRequest    time.Time `orm:"type(datetime)" form:"dateRequest,02-01-2006 15:04:05 -0700"`
	DateEstimated  time.Time `orm:"null;type(datetime)" form:"dateEstimated,02-01-2006 15:04:05 -0700"`
	DateResolution time.Time `orm:"null;type(datetime)" form:"dateResolution,02-01-2006 15:04:05 -0700"`
	Priority       int       `orm:"type(number)" form:"priority" valid:"Range(1,4)"`
	ConfirmUser    int       `orm:"null;type(number)" form:"confirmUser" valid:"Range(0,3)"`
	User           *User     `orm:"rel(fk);null; on_delete (do_nothing)"`
}

//renvoie de la structure à travers une fonction via Incident.TableIncident
func (i *Incident) TableName() string {
	return "incident"
}
