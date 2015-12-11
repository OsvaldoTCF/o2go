package models

// supplier

import (
	"github.com/gernest/utron"
	"github.com/jinzhu/gorm"
	//	"time"
)

type Supplier struct {
	gorm.Model
	Name        string          `schema:"name" sql:"type:varchar(100);not null"`
	PhoneNumber string          `schema:"-"`
	Note        string          `schema:"note" sql:"type:varchar(250);"`
	Emails      []EmailSupplier `schema:"-"`
}

func init() {
	utron.RegisterModels(&Supplier{})
}
