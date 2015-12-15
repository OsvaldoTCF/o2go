package controllers

import (
	"encoding/json"
	"github.com/OsvaldoTCF/order2go/models"
	"github.com/gernest/utron"
	"net/http"
	"strconv"
)

type EmailSponsorAction struct {
	*utron.BaseController
	Routes []string
}

func (a *EmailSponsorAction) Get() {
	var emails []*models.EmailSponsor

	vars, ok1 := a.Ctx.Params["id"]
	var id int
	var err1 error = nil
	if ok1 {
		id, err1 = strconv.Atoi(vars)
	}

	a.Ctx.Request().ParseForm()
	par, ok2 := a.Ctx.Request().Form["sponsor"]
	var sup int
	var err2 error = nil
	if ok2 {
		sup, err2 = strconv.Atoi(par[0])
	}

	if (err1 == nil) && ok1 {
		a.Ctx.DB.Where("id = ?", id).Find(&emails)
		var s models.Sponsor

		a.Ctx.DB.Where("id = ?", emails[0].SponsorId).Find(&s)

		data := struct {
			Length  int
			Sponsor string
			Emails  []*models.EmailSponsor
		}{
			len(emails),
			s.Name,
			emails,
		}

		a.RenderJSON(data, http.StatusOK)
	} else if (err2 == nil) && ok2 {
		a.Ctx.DB.Where("sponsor_id = ?", sup).Order("id").Limit(20).Find(&emails)
		var s models.Sponsor

		a.Ctx.DB.Where("id = ?", sup).Find(&s)

		data := struct {
			Length  int
			Sponsor string
			Emails  []*models.EmailSponsor
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
			Emails []*models.EmailSponsor
		}{
			len(emails),
			emails,
		}

		a.RenderJSON(data, http.StatusOK)
	}
}

func (a *EmailSponsorAction) Post() {
	a.Ctx.Request().ParseForm()
	dec := json.NewDecoder(a.Ctx.Request().Body)

	var m map[string]interface{}
	dec.Decode(&m)

	eml := new(models.EmailSponsor)
	eml.Email, _ = m["email"].(string)

	i, _ := strconv.Atoi(m["sponsor"].(string))
	eml.SponsorId = uint(i)

	a.Ctx.DB.Create(&eml)

	a.RenderJSON(eml, http.StatusOK)
}

func (a *EmailSponsorAction) Put() {
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

		eml := models.EmailSponsor{}

		a.Ctx.DB.Where("id = ?", id).Find(&eml)

		eml.Email, _ = m["email"].(string)

		a.Ctx.DB.Save(&eml)

		a.RenderJSON(eml, http.StatusOK)
	}
}

func (a *EmailSponsorAction) Delete() {
	vars, ok := a.Ctx.Params["id"]
	var id int
	var err error = nil
	if ok {
		id, err = strconv.Atoi(vars)
	}

	if (err == nil) && ok {
		a.Ctx.DB.Exec("delete from email_sponsors where id = ?", id)
	}
	a.Ctx.Set(http.StatusNoContent)
}

func NewEmailSponsorAction() *EmailSponsorAction {
	return &EmailSponsorAction{
		Routes: []string{
			"get;/emailpatrocinador/{id};Get",
			"get;/emailpatrocinador;Get",
			"post;/emailpatrocinador;Post",
			"put;/emailpatrocinador/{id};Put",
			"delete;/emailpatrocinador/{id};Delete",
		},
	}
}

func init() {
	utron.RegisterController(NewEmailSupplierAction())
}
