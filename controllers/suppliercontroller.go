package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/OsvaldoTCF/order2go/models"
	_ "github.com/OsvaldoTCF/order2go/utils"
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

	//	s := models.Supplier{}
	//	s.Name = "KJ Mouses"
	//	s.PhoneNumber = "9 8765-4321"
	//	s.Note = "Trustworthy"
	//	a.Ctx.DB.Create(&s)

	if (err == nil) && ok {
		a.Ctx.DB.Where("id = ?", id).Find(&suppliers)
		for i, _ := range suppliers {
			a.Ctx.DB.Where("supplier_id = ?", suppliers[i].ID).Order("id").Find(&emails)
			suppliers[i].Emails = emails
			suppliers[i].SetEmailsLength()
		}
	} else {
		a.Ctx.DB.Order("id").Limit(20).Find(&suppliers)
		for i, _ := range suppliers {
			a.Ctx.DB.Where("supplier_id = ?", suppliers[i].ID).Order("id").Find(&emails)
			suppliers[i].Emails = emails
			suppliers[i].SetEmailsLength()
		}
	}

	data := struct {
		Length    int
		Suppliers []*models.Supplier
	}{
		len(suppliers),
		suppliers,
	}

	rtn, _ := json.Marshal(data)

	a.Ctx.Set(http.StatusOK)
	a.Ctx.Write(rtn)

	fmt.Println(string(rtn[:]))
}

func (a *SupplierAction) Post() {
	a.Ctx.Request().ParseForm()
	dec := json.NewDecoder(a.Ctx.Request().Body)

	var m map[string]interface{}
	dec.Decode(&m)
	//	ems, _ := m["emails"].([]interface{})

	sup := new(models.Supplier)
	sup.Name, _ = m["name"].(string)
	sup.Note, _ = m["note"].(string)
	sup.PhoneNumber, _ = m["phone"].(string)
	//	sup.Emails = []models.EmailSupplier{}

	//	for _, val := range ems {
	//		v := val.(map[string]interface{})
	//		s, _ := v["email"]
	//		aux := models.EmailSupplier{
	//			Email: s.(string),
	//		}
	//		sup.Emails = append(sup.Emails, aux)
	//	}

	a.Ctx.DB.Create(&sup)

	rtn, _ := json.Marshal(sup)

	a.Ctx.Set(http.StatusCreated)
	a.Ctx.Write(rtn)
	fmt.Println(string(rtn[:]))
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
		ems, _ := m["emails"].([]interface{})

		sup := models.Supplier{}

		a.Ctx.DB.Where("id = ?", id).Find(&sup)

		sup.Name, _ = m["name"].(string)
		sup.Note, _ = m["note"].(string)
		sup.PhoneNumber, _ = m["phone"].(string)
		sup.Emails = []models.EmailSupplier{}

		for _, val := range ems {
			v := val.(map[string]interface{})
			s, _ := v["email"]
			aux := models.EmailSupplier{
				Email: s.(string),
			}
			sup.Emails = append(sup.Emails, aux)
		}

		a.Ctx.DB.Save(&sup)

		rtn, _ := json.Marshal(sup)
		a.Ctx.Write(rtn)
		fmt.Println(string(rtn[:]))
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
