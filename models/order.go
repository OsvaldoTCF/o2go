package models
// Order

import (
	"github.com/gernest/utron"
	"time"
)

type Order struct {
	ID        uint
	SupplierId uint      `sql:"not null"`
	Valor     float32   `sql:"type:numeric(18,4);not null"`
	Vencto    time.Time `sql:"not null"`
	Data   time.Time `sql:"DEFAULT:'1970-01-01 00:00:00-03'"`
	OrderItems 
}

func (d *Order) SetDate() {
	d.Date = d.Datapag.Format("02/01/2006 15:04:05")
}

func init() {
	utron.RegisterModels(&Order{})
}
