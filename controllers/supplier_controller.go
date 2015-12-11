package controllers

import (
	"encoding/json"
	"github.com/OsvaldoTCF/order2go/models"
	"github.com/gernest/utron"
	"net/http"
	"strconv"
)

type SupplierAction struct {
	*utron.BaseController
	Routes []string
}

func (a *SupplierAction) Get() {
	var suppliers []*models.Supplier
	var emails []models.EmailSupplier
	vars, ok := a.Ctx.Params["id"]
	var id int
	var err error = nil
	if ok {
		id, err = strconv.Atoi(vars)
	}

	if (err == nil) && ok {
		a.Ctx.DB.Where("id = ?", id).Find(&suppliers)
		for i, _ := range suppliers {
			a.Ctx.DB.Where("supplier_id = ?", suppliers[i].ID).Order("id").Find(&emails)
			suppliers[i].Emails = emails
		}
	} else {
		a.Ctx.DB.Order("id").Limit(20).Find(&suppliers)
		for i, _ := range suppliers {
			a.Ctx.DB.Where("supplier_id = ?", suppliers[i].ID).Order("id").Find(&emails)
			suppliers[i].Emails = emails
		}
	}

	data := struct {
		Length    int
		Suppliers []*models.Supplier
	}{
		len(suppliers),
		suppliers,
	}

	a.RenderJSON(data, http.StatusOK)
}

func (a *SupplierAction) Post() {
	a.Ctx.Request().ParseForm()
	dec := json.NewDecoder(a.Ctx.Request().Body)

	var m map[string]interface{}
	dec.Decode(&m)

	sup := new(models.Supplier)
	sup.Name, _ = m["name"].(string)
	sup.Note, _ = m["note"].(string)
	sup.PhoneNumber, _ = m["phone"].(string)

	a.Ctx.DB.Create(&sup)

	a.RenderJSON(sup, http.StatusOK)
}

func (a *SupplierAction) Put() {
	vars, ok := a.Ctx.Params["id"]
	var id int
	var err error = nil
	if ok {
		id, err = strconv.Atoi(vars)
	}

	if (err == nil) && ok {

		a.Ctx.Request().ParseForm()
		dec := json.NewDecoder(a.Ctx.Request().Body)

		var m map[string]interface{}
		dec.Decode(&m)

		sup := models.Supplier{}

		a.Ctx.DB.Where("id = ?", id).Find(&sup)

		sup.Name, _ = m["name"].(string)
		sup.Note, _ = m["note"].(string)
		sup.PhoneNumber, _ = m["phone"].(string)

		a.Ctx.DB.Save(&sup)

		a.RenderJSON(sup, http.StatusOK)
	}
}

func (a *SupplierAction) Delete() {
	vars, ok := a.Ctx.Params["id"]
	var id int
	var err error = nil
	if ok {
		id, err = strconv.Atoi(vars)
	}

	if (err == nil) && ok {
		a.Ctx.DB.Exec("delete from suppliers where id = ?", id)
	}
	a.Ctx.Set(http.StatusNoContent)
}

func NewSupplierAction() *SupplierAction {
	return &SupplierAction{
		Routes: []string{
			"get;/fornecedores/{id};Get",
			"get;/fornecedores;Get",
			"post;/fornecedores;Post",
			"put;/fornecedores/{id};Put",
			"delete;/fornecedores/{id};Delete",
		},
	}
}

func init() {
	utron.RegisterController(NewSupplierAction())
}
