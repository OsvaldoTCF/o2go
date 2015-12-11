package controllers

import (
	"encoding/json"
	"github.com/OsvaldoTCF/order2go/models"
	"github.com/gernest/utron"
	"net/http"
	"strconv"
)

type EmailSupplierAction struct {
	*utron.BaseController
	Routes []string
}

func (a *EmailSupplierAction) Get() {
	var emails []*models.EmailSupplier

	vars, ok1 := a.Ctx.Params["id"]
	var id int
	var err1 error = nil
	if ok1 {
		id, err1 = strconv.Atoi(vars)
	}

	a.Ctx.Request().ParseForm()
	par, ok2 := a.Ctx.Request().Form["supplier"]
	var sup int
	var err2 error = nil
	if ok2 {
		sup, err2 = strconv.Atoi(par[0])
	}

	if (err1 == nil) && ok1 {
		a.Ctx.DB.Where("id = ?", id).Find(&emails)
		var s models.Supplier

		a.Ctx.DB.Where("id = ?", emails[0].SupplierID).Find(&s)

		data := struct {
			Length   int
			Supplier string
			Emails   []*models.EmailSupplier
		}{
			len(emails),
			s.Name,
			emails,
		}

		a.RenderJSON(data, http.StatusOK)
	} else if (err2 == nil) && ok2 {
		a.Ctx.DB.Where("supplier_id = ?", sup).Order("id").Limit(20).Find(&emails)
		var s models.Supplier

		a.Ctx.DB.Where("id = ?", sup).Find(&s)

		data := struct {
			Length   int
			Supplier string
			Emails   []*models.EmailSupplier
		}{
			len(emails),
			s.Name,
			emails,
		}

		a.RenderJSON(data, http.StatusOK)
	} else {
		a.Ctx.DB.Order("id").Limit(20).Find(&emails)

		data := struct {
			Length int
			Emails []*models.EmailSupplier
		}{
			len(emails),
			emails,
		}

		a.RenderJSON(data, http.StatusOK)
	}
}

func (a *EmailSupplierAction) Post() {
	a.Ctx.Request().ParseForm()
	dec := json.NewDecoder(a.Ctx.Request().Body)

	var m map[string]interface{}
	dec.Decode(&m)

	eml := new(models.EmailSupplier)
	eml.Email, _ = m["email"].(string)

	i, _ := strconv.Atoi(m["supplier"].(string))
	eml.SupplierID = uint(i)

	a.Ctx.DB.Create(&eml)

	a.RenderJSON(eml, http.StatusOK)
}

func (a *EmailSupplierAction) Put() {
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

		eml := models.EmailSupplier{}

		a.Ctx.DB.Where("id = ?", id).Find(&eml)

		eml.Email, _ = m["email"].(string)

		a.Ctx.DB.Save(&eml)

		a.RenderJSON(eml, http.StatusOK)
	}
}

func (a *EmailSupplierAction) Delete() {
	vars, ok := a.Ctx.Params["id"]
	var id int
	var err error = nil
	if ok {
		id, err = strconv.Atoi(vars)
	}

	if (err == nil) && ok {
		a.Ctx.DB.Exec("delete from email_suppliers where id = ?", id)
	}
	a.Ctx.Set(http.StatusNoContent)
}

func NewEmailSupplierAction() *EmailSupplierAction {
	return &EmailSupplierAction{
		Routes: []string{
			"get;/emailfornecedor/{id};Get",
			"get;/emailfornecedor;Get",
			"post;/emailfornecedor;Post",
			"put;/emailfornecedor/{id};Put",
			"delete;/emailfornecedor/{id};Delete",
		},
	}
}

func init() {
	utron.RegisterController(NewEmailSupplierAction())
}
