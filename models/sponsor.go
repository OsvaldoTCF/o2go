package models
// sponsor

import (
	"github.com/gernest/utron"
	//"time"
)

type Sponsor struct {
	ID        uint  `schema: "-"`
	Name		  string `schema: "name"` `sql:"type:varchar(100);not null"`
	PhoneExtension  uint  `schema: "phonextension"` `sql:not null"`
	Note       string  `schema: "note"` `sql:"type:varchar(250);DEFAULT:''"`
    Emails    []Email 
	CreatedAt time.Time `schema:"-"`
    UpdatedAt time.Time `schema:"-"`
    DeletedAt *time.Time `schema:"-"`
}

func init() {
	utron.RegisterModels(&Sponsor{})
}
