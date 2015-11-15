package models
// Order

import (
	"github.com/gernest/utron"
	"time"
)

type Order struct {
	ID        uint  `schema: "-"`
	SupplierId uint   `schema: "-"`  `sql:"not null"`
	SponsorId uint  `schema: "-"` `sql:"not null"`
	Valor     float32   `schema: "valor"` `sql:"type:numeric(7,2);not null"`
	DueDateBase    time.Time `schema: "duedatebase"` `sql:"not null;DEFAULT:current_timestamp"`
	Date   time.Time `schema: "date"` `sql:"DEFAULT:'1970-01-01 00:00:00-03'"`
	OrderItems	[]OrderItem 
}

func (o *Order) SetDate() {
	o.Date = o.Date.Format("02/01/2006 15:04:05")
}

func init() {
	utron.RegisterModels(&Order{})
}
