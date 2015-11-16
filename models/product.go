package models
// product

import (
	"github.com/gernest/utron"
	"time"
)

type Product struct {
	ID        uint `schema: "-"`
	Name string    `schema:"name";sql:"type:varchar(100);not null"`
	Supplier_ID	uint   `schema: "-";sql:"not null"`
	Price     float32   `schema:"price";sql:"type:numeric(18,4);not null"`
	Unit		string `schema:"unit";sql:size:3`
	CreatedAt time.Time `schema:"-"`
    UpdatedAt time.Time `schema:"-"`
    DeletedAt time.Time `schema:"-"`
}


func init() {
	utron.RegisterModels(&Product{})
}
