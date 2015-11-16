package models
// product

import (
	"github.com/gernest/utron"
//	"time"
)

type Product struct {
	Model	`schema:"-"`
	Name string    `schema:"name";sql:"type:varchar(100);not null"`
	Supplier		Supplier
	SupplierId	uint   `schema: "-";sql:"not null"`
	Price     float32   `schema:"price";sql:"type:numeric(18,4);not null"`
	Unit		string `schema:"unit";sql:size:3`
}


func init() {
	utron.RegisterModels(&Product{})
}
