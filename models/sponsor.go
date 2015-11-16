package models
// sponsor

import (
	"github.com/gernest/utron"
	//"time"
)

type Sponsor struct {
	Gmodel  `schema: "-"`
	Name		  string `schema: "name";sql:"type:varchar(100);not null"`
	PhoneExtension  uint  `schema: "phonextension";sql:not null"`
	Note       string  `schema: "note";sql:"type:varchar(250);DEFAULT:''"`
    Emails    []Email 
}

func init() {
	utron.RegisterModels(&Sponsor{})
}
