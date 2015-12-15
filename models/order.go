package models

// Order

import (
	"github.com/gernest/utron"
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Supplier    Supplier    `schema:"-"`
	SupplierId  uint        `schema:"-" sql:"type:integer REFERENCES Suppliers(id)"`
	SponsorId   uint        `schema:"-" sql:"type:integer REFERENCES Sponsors(id)"`
	Sponsor     Sponsor     `schema:"-"`
	Value       float32     `schema:"value" sql:"type:numeric(7,2);not null"`
	DueDateBase time.Time   `schema: "duedatebase" sql:"not null;DEFAULT:current_timestamp"`
	OrderItems  []OrderItem `schema:"-"`
}

//func (o *Order) SetDate() {
//	o.Date = o.Date.Format("02/01/2006 15:04:05")
//}

func init() {
	utron.RegisterModels(&Order{})
}
