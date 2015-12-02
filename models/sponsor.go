package models

// sponsor

import (
	_ "github.com/gernest/utron"
	"github.com/jinzhu/gorm"
	//"time"
)

type Sponsor struct {
	gorm.Model
	Name     string         `schema:"name" sql:"type:varchar(100);not null"`
	PhoneExt uint           `schema:"phonext" sql:not null"` // Phone Extension
	Note     string         `schema:"note" sql:"type:varchar(250);DEFAULT:'Caution'"`
	Emails   []EmailSponsor `schema:"-"`
}

//func init() {
//	utron.RegisterModels(&Sponsor{})
//}
