package controllers

import (
	"github.com/OsvaldoTCF/order2go/utils"
	"github.com/gernest/utron"
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

type GeneralAction struct {
	*utron.BaseController
	Routes []string
}

type Action interface {
	Get()
	Post()
	Put()
	Delete()
}

func (a *GeneralAction) WrongMethod() {
	a.Ctx.Set(http.StatusMethodNotAllowed)
	a.Ctx.Write(utils.BuildReturnMessage(utils.OPERATION_NOT_ALLOWED))
}

func (a *GeneralAction) Home() {
	a.Ctx.Template = "index"
	a.Ctx.HTML()
}

func NewGeneralAction() *GeneralAction {
	return &GeneralAction{
		Routes: []string{
			"post,put,delete;/;WrongMethod",
			"get;/;Home",
		},
	}
}

func init() {
	utron.RegisterController(NewGeneralAction())
}
