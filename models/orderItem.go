package models
// Order Item

import (
	"github.com/gernest/utron"
	//"time"
)

type OrderItem struct {
	OrderId uint `schema:"-" sql:"type:integer REFERENCES orders(id)"`
	ProductId  uint `schema:"-" sql:"type:integer REFERENCES products(id)"`
	Qty  float32 `schema:"qty" sql:"type:numeric(7,3)"`
	TotalPrice  float32 `schema:"totalprice" sql:"type:numeric(7,2)"`

}


func init() {
	utron.RegisterModels(&OrderItem{})
}
