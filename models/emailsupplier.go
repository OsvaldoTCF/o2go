package models
// EmailSupplier

import (
	"github.com/gernest/utron"
//	"time"
)

type EmailSupplier struct {
	Model	`schema:"-"`
    SupplierID  uint `schema:"-" sql:"type:integer REFERENCES Suppliers(id)"`
    Email   string  `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
}

func init() {
	utron.RegisterModels(&EmailSupplier{})
}
