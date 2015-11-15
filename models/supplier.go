package models
// supplier

import (
	"github.com/gernest/utron"
	//"time"
)

type Supplier struct {
	ID        uint
	Name		  string     `sql:"type:varchar(100);not null"`
	PhoneNumber  string  `sql:"type:varchar(11);not null"`
	Obs       string     `sql:"type:varchar(250);DEFAULT:''"`
    Emails    []Email 
	CreatedAt time.Time `schema:"-"`
    UpdatedAt time.Time `schema:"-"`
    DeletedAt time.Time `schema:"-"`
}

func init() {
	utron.RegisterModels(&Supplier{})
}
