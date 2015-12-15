package controllers

import (
	"encoding/json"
	"github.com/OsvaldoTCF/order2go/models"
	"github.com/gernest/utron"
	"net/http"
	"strconv"
)

type SponsorAction struct {
	*utron.BaseController
	Routes []string
}

func (a *SponsorAction) Get() {
	var sponsors []*models.Sponsor
	var emails []models.EmailSponsor
	vars, ok := a.Ctx.Params["id"]
	var id int
	var err error = nil
	if ok {
		id, err = strconv.Atoi(vars)
	}

	if (err == nil) && ok {
		a.Ctx.DB.Where("id = ?", id).Find(&sponsors)
		for i, _ := range sponsors {
			a.Ctx.DB.Where("Sponsor_id = ?", sponsors[i].ID).Order("id").Find(&emails)
			sponsors[i].Emails = emails
		}
	} else {
		a.Ctx.DB.Order("id").Limit(20).Find(&sponsors)
		for i, _ := range sponsors {
			a.Ctx.DB.Where("sponsor_id = ?", sponsors[i].ID).Order("id").Find(&emails)
			sponsors[i].Emails = emails
		}
	}

	data := struct {
		Length   int
		Sponsors []*models.Sponsor
	}{
		len(sponsors),
		sponsors,
	}

	a.RenderJSON(data, http.StatusOK)
}

func (a *SponsorAction) Post() {
	a.Ctx.Request().ParseForm()
	dec := json.NewDecoder(a.Ctx.Request().Body)

	var m map[string]interface{}
	dec.Decode(&m)

	sup := new(models.Sponsor)
	sup.Name, _ = m["name"].(string)
	sup.Note, _ = m["note"].(string)
	px, _ := m["phone"].(string)
	pxt, _ := strconv.Atoi(px)
	sup.PhoneExt = uint(pxt)

	a.Ctx.DB.Create(&sup)

	a.RenderJSON(sup, http.StatusOK)
}

func (a *SponsorAction) Put() {
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

		sup := models.Sponsor{}

		a.Ctx.DB.Where("id = ?", id).Find(&sup)

		sup.Name, _ = m["name"].(string)
		sup.Note, _ = m["note"].(string)
		px, _ := m["phone"].(string)
		pxt, _ := strconv.Atoi(px)
		sup.PhoneExt = uint(pxt)

		a.Ctx.DB.Save(&sup)

		a.RenderJSON(sup, http.StatusOK)
	}
}

func (a *SponsorAction) Delete() {
	vars, ok := a.Ctx.Params["id"]
	var id int
	var err error = nil
	if ok {
		id, err = strconv.Atoi(vars)
	}

	if (err == nil) && ok {
		a.Ctx.DB.Exec("delete from sponsors where id = ?", id)
	}
	a.Ctx.Set(http.StatusNoContent)
}

func NewSponsorAction() *SponsorAction {
	return &SponsorAction{
		Routes: []string{
			"get;/patrocinadores/{id};Get",
			"get;/patrocinadores;Get",
			"post;/patrocinadores;Post",
			"put;/patrocinadores/{id};Put",
			"delete;/patrocinadores/{id};Delete",
		},
	}
}

func init() {
	utron.RegisterController(NewSponsorAction())
}
