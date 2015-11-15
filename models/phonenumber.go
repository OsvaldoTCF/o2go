package models
// email

import (
	"github.com/gernest/utron"
	//"time"
)

type Email struct {
    ID      int
    SupplierID  int     `sql:"index"` // Foreign key (belongs to), tag `index` will create index for this field when using AutoMigrate
    Email   string  `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
    Subscribed bool
	CreatedAt time.Time `schema:"-"`
    UpdatedAt time.Time `schema:"-"`
    DeletedAt time.Time `schema:"-"`
}

func init() {
	utron.RegisterModels(&Email{})
}
