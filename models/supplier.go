package models
// supplier

import (
	"github.com/gernest/utron"
	"time"
)

type Supplier struct {
	ID        uint  `schema: "-"`
	Name		  string `schema: "name";sql:"type:varchar(100);not null"`
	PhoneNumbers  []PhoneNumber  `schema: "-"`
	Note       string  `schema: "note";sql:"type:varchar(250);"`
    Emails    []Email  
	CreatedAt time.Time `schema:"-"`
    UpdatedAt time.Time `schema:"-"`
    DeletedAt *time.Time `schema:"-"`
}

func init() {
	utron.RegisterModels(&Supplier{})
}
