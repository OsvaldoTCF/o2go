// debito
package models

import (
	"github.com/gernest/utron"
	"time"
)

type Debito struct {
	ID        uint
	Documento string    `sql:"type:varchar(20);not null"`
	Empid     uint      `sql:"not null"`
	Clid      uint      `sql:"not null"`
	Cartao    string    `sql:"type:varchar(16);not null"`
	Nome      string    `sql:"type:varchar(255);not null"`
	Telefone  string    `sql:"type:varchar(11);not null"`
	CpfCnpj   string    `sql:"type:varchar(14);not null"`
	Endereco  string    `sql:"type:varchar(250);not null"`
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

func (d *Debito) SetDate() {
	d.Date = d.Datapag.Format("02/01/2006 15:04:05")
}

func init() {
	utron.RegisterModels(&Debito{})
}
