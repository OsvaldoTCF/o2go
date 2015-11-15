package controllers

import (
	"encoding/json"
	_ "fmt"
	"github.com/gernest/utron"
	_ "naraka/models"
	"naraka/utils"
	"net/http"
	"strconv"
	"time"
)

type OpeningAction struct {
	*utron.BaseController
	Routes []string
}

func (a *OpeningAction) WrongMethod() {
	a.HTML(http.StatusMethodNotAllowed)
	a.Ctx.Template = "error"
	a.Ctx.Data["Message"] = utils.BuildReturnMessage(utils.OPERATION_NOT_ALLOWED)
}

func (a *OpeningAction) Opening() {
	params := a.Ctx.Request().URL.Query()

	if len(params) != 1 {
		a.HTML(http.StatusNotFound)
		a.Ctx.Template = "error"
		a.Ctx.Data["Message"] = utils.BuildReturnMessage(utils.RESOURCE_NOT_FOUND)
	} else if term, ok := params["terminal"]; ok == false {
		a.HTML(http.StatusNotFound)
		a.Ctx.Template = "error"
		a.Ctx.Data["Message"] = utils.BuildReturnMessage(utils.RESOURCE_NOT_FOUND)
	} else {
		if ter, err := strconv.Atoi(term[0]); err != nil {
			a.HTML(http.StatusBadRequest)
			a.Ctx.Template = "error"
			a.Ctx.Data["Message"] = utils.BuildReturnMessage(utils.INVALIDE_PARAMETER_TYPE)
		} else {
			contains := func(sl []int, nm int) bool {
				for _, a := range sl {
					if a == nm {
						return true
					}
				}
				return false
			}

			if !contains(utils.TERMS, ter) {
				a.HTML(http.StatusBadRequest)
				a.Ctx.Template = "error"
				a.Ctx.Data["Message"] = utils.BuildReturnMessage(utils.INVALIDE_TERMINAL_VALUE)
			} else {
				d := time.Now()
				wtc := utils.GetWatcherInstance()
				if ok, _ := wtc.IsOpen(ter); !ok {
					wtc.OpenTerminal(ter, d)
				}

				ctt := struct {
					Name    string `json:"nome_impressora"`
					Address string `json:"endereco_impressora"`
					Stamp   string `json:"horario"`
				}{
					utils.PRINTERS[ter],
					utils.ADDRESSES[ter],
					d.Format("02/01/2006 15:04:05"),
				}

				rtn := struct {
					Content interface{} `json:"abertura"`
				}{
					ctt,
				}

				b, _ := json.Marshal(rtn)

				a.HTML(http.StatusOK)
				a.Ctx.Template = "message"
				a.Ctx.Data["Message"] = string(b[:])
			}
		}
	}
}

func NewOAction() *OpeningAction {
	return &OpeningAction{
		Routes: []string{
			"get;/abertura;Opening",
			"post,put,delete,patch;/abertura;WrongMethod",
		},
	}
}

func init() {
	utron.RegisterController(NewOAction())
}
