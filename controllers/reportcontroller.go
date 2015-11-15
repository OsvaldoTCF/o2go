package controllers

import (
	"github.com/gernest/utron"
	"github.com/gorilla/schema"
	"naraka/models"
	"naraka/utils"
	"net/http"
	"strconv"
)

var decoder = schema.NewDecoder()

type ReportAction struct {
	*utron.BaseController
	Routes []string
}

func (a *ReportAction) WrongMethod() {
	a.HTML(http.StatusMethodNotAllowed)
	a.Ctx.Template = "error"
	a.Ctx.Data["Message"] = utils.BuildReturnMessage(utils.OPERATION_NOT_ALLOWED)
}

func (a *ReportAction) Report() {
	debitos := []*models.Debito{}
	params := a.Ctx.Request().URL.Query()

	if len(params) == 0 {
		a.Ctx.DB.Where("pago = ? and datapag between ? and ?", "true", "'TODAY'", "now()").Order("datapag").Find(&debitos)
		a.Ctx.Data["List"] = debitos
		a.Ctx.Template = "index"
		a.HTML(http.StatusOK)
	} else {
		if term, ok := params["terminal"]; ok == true {
			ter, _ := strconv.Atoi(term[0])
			a.Ctx.DB.Where("pago = ? and terminal = ? and datapag between ? and ?", "true", ter, "'TODAY'", "now()").Order("datapag").Find(&debitos)
			for j, _ := range debitos {
				debitos[j].SetDate()
			}
			qty, ttl := TermTotal(debitos)
			data := struct {
				Trm, Qtd int
				Tot      float32
			}{
				ter,
				qty,
				ttl,
			}
			a.Ctx.Data["Data"] = data
			a.Ctx.Data["List"] = debitos
			a.Ctx.Template = "report"
			a.HTML(http.StatusOK)
		} else {
			a.Ctx.Redirect("/relatorio", http.StatusOK)
		}
	}
}

func TermTotal(d []*models.Debito) (count int, sum float32) {
	count = len(d)
	sum = 0
	if count > 0 {
		for _, val := range d {
			sum += val.Valpag
		}
	}
	return
}

func NewAction() *ReportAction {
	return &ReportAction{
		Routes: []string{
			"get;/relatorio;Report",
			"post,put,delete,patch;/relatorio;WrongMethod",
		},
	}
}

func init() {
	utron.RegisterController(NewAction())
}
