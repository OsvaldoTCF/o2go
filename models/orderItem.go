package models
// Order Item

import (
	"github.com/gernest/utron"
	"time"
)

type OrderItem struct {
	ID        uint
	Documento string    `sql:"type:varchar(20);not null"`
	Empid     uint      `sql:"not null"`
	Cidade    string    `sql:"type:varchar(30);not null"`
	Obs       string    `sql:"type:varchar(30);DEFAULT:''"`
	Valor     float32   `sql:"type:numeric(18,4);not null"`
	Valpag    float32   `sql:"type:numeric(18,4)"`
	Vencto    time.Time `sql:"not null"`
	Datapag   time.Time `sql:"DEFAULT:'1970-01-01 00:00:00-03'"`
	Lote      uint      `sql:"not null"`
	Terminal  uint      `sql:"DEFAULT:0"`
	Pago      bool      `sql:"DEFAULT:false"`
	Date      string    `sql:"-"`
}


func init() {
	utron.RegisterModels(&OrderItem{})
}
