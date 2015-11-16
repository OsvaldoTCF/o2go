package models
// email

import (
	"github.com/gernest/utron"
	"time"
)

type Email struct {
    ID      uint
    SupplierID  uint `schema:"-" sql:"type:integer REFERENCES Suppliers(id)"`
    Email   string  `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
	CreatedAt time.Time `schema:"-"`
    UpdatedAt time.Time `schema:"-"`
    DeletedAt time.Time `schema:"-"`
}

func init() {
	utron.RegisterModels(&Email{})
}
