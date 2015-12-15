package models

// EmailSponsor

import (
	"github.com/gernest/utron"
	"github.com/jinzhu/gorm"
	//	"time"
)

type EmailSponsor struct {
	gorm.Model
	SponsorId uint   `schema:"-" gorm:"primary_key" sql:"type:integer REFERENCES sponsors(id)"`
	Email     string `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
}

func init() {
	utron.RegisterModels(&EmailSponsor{})
}
