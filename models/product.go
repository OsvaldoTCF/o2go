package models

// product

import (
	"github.com/gernest/utron"
	"github.com/jinzhu/gorm"
	//	"time"
)

type Product struct {
	gorm.Model
	Name       string   `schema:"name" sql:"type:varchar(100);not null"`
	Supplier   Supplier `schema:"-"`
	SupplierId uint     `sql:"type:integer REFERENCES Suppliers(id)"`
	Price      float32  `schema:"price" sql:"type:numeric(7,2);not null"`
	Unit       string   `schema:"unit" sql:size:3`
}

func init() {
	utron.RegisterModels(&Product{})
}
