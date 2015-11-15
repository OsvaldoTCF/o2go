package models
// product

import (
	"github.com/gernest/utron"
	//"time"
)

type Product struct {
	ID        uint
	Name string    `sql:"type:varchar(100);not null"`
	Supplier_ID	uint   `sql:"not null"`
	Price     float32   `sql:"type:numeric(18,4);not null"`
	CreatedAt time.Time `schema:"-"`
    UpdatedAt time.Time `schema:"-"`
    DeletedAt time.Time `schema:"-"`
}


func init() {
	utron.RegisterModels(&Product{})
}
