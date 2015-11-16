package models
// sponsor

import (
	"github.com/gernest/utron"
	//"time"
)

type Sponsor struct {
	Model	`schema:"-"`
	Name		  string `schema:"name" sql:"type:varchar(100);not null"`
	PhoneExt  uint  `schema:"phonext" sql:not null"` // Phone Extension
	Note       string  `schema:"note" sql:"type:varchar(250);DEFAULT:'Caution'"`
    Emails    []EmailSponsor `schema:"-"`
}

func init() {
	utron.RegisterModels(&Sponsor{})
}
