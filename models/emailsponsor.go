package models
// EmailSponsor

import (
	"github.com/gernest/utron"
//	"time"
)

type EmailSponsor struct {
	Model	`schema:"-"`
    SponsorId  uint `schema:"-" gorm:"primary_key" sql:"type:integer REFERENCES Sponsors(id)"`
    Email   string  `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
}

func init() {
	utron.RegisterModels(&EmailSponsor{})
}
