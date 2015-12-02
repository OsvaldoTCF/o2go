package models

// EmailSupplier

import (
	"github.com/gernest/utron"
	"github.com/jinzhu/gorm"
	//	"time"
)

type EmailSupplier struct {
	gorm.Model
	SupplierID uint   `schema:"-" gorm:"primary_key" sql:"type:integer REFERENCES Suppliers(id)"`
	Email      string `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
}

func init() {
	utron.RegisterModels(&EmailSupplier{})
}
