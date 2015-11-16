package models
// Order

import (
	"github.com/gernest/utron"
	"time"
)

type Order struct {
	Model	`schema:"-"`
	Supplier		Supplier
	SupplierId uint `schema:"-" sql:"type:integer REFERENCES Suppliers(id)"`
	SponsorId uint  `schema:"-" sql:"type:integer REFERENCES Sponsors(id)"`
	Sponsor  Sponsor
	Valor     float32   `schema:"valor" sql:"type:numeric(7,2);not null"`
	DueDateBase    time.Time `schema: "vencto" sql:"not null;DEFAULT:current_timestamp"`
	OrderItems	[]OrderItem 
}

//func (o *Order) SetDate() {
//	o.Date = o.Date.Format("02/01/2006 15:04:05")
//}

func init() {
	utron.RegisterModels(&Order{})
}
