package utils

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

var lote int

func Lote() int {
	lote += 1
	return lote
}

func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func hasElement(vs []string, t string) bool {
	return index(vs, t) >= 0
}

func ValidateTel(tel string) (ok bool) {
	ok = true
	s := strings.TrimSpace(tel)
	s = strings.Replace(s, " ", "", -1)
	if len(s) > 11 {
		ok = false
	}
	if hasElement(INVALIDES, s) {
		ok = false
	}
	_, err := strconv.Atoi(s)
	if err != nil {
		ok = false
	}
	return
}

func ValidateCard(card string) (ok bool) {
	ok = true
	s := strings.TrimSpace(card)
	s = strings.Replace(s, " ", "", -1)
	if len(s) != 16 {
		ok = false
	}
	_, err := strconv.Atoi(s)
	if err != nil {
		ok = false
	}
	return
}

func validateCpf(cpf string) (ok bool) {
	var s, z, p, i, r int
	var ten, eleven rune

	ok = true
	if hasElement(INVALIDES, cpf) {
		ok = false
	}

	s = 0
	p = 10

	for i = 0; i <= 8; i++ {
		z = int(cpf[i] - 48)
		s += z * p
		p--
	}

	r = 11 - (s % 11)
	if (r == 10) || (r == 11) {
		ten = '0'
	} else {
		ten = rune(r + 48)
	}

	s = 0
	p = 11

	for i = 0; i <= 9; i++ {
		z = int(cpf[i] - 48)
		s += z * p
		p--
	}

	r = 11 - (s % 11)
	if (r == 10) || (r == 11) {
		eleven = '0'
	} else {
		eleven = rune(r + 48)
	}

	if (ten != rune(cpf[9])) || (eleven != rune(cpf[10])) {
		ok = false
	}
	return
}

func validateCnpj(cnpj string) (ok bool) {
	var s, z, p, i, r int
	var thirteen, fourteen rune

	ok = true
	if hasElement(INVALIDES, cnpj) {
		ok = false
	}

	s = 0
	p = 2

	for i = 11; i >= 0; i-- {
		z = int(cnpj[i] - 48)
		s += z * p
		p++
		if p == 10 {
			p = 2
		}
	}

	r = s % 11
	if (r == 0) || (r == 1) {
		thirteen = '0'
	} else {
		thirteen = rune((11 - r) + 48)
	}

	s = 0
	p = 2

	for i = 12; i >= 0; i-- {
		z = int(cnpj[i] - 48)
		s += z * p
		p++
		if p == 10 {
			p = 2
		}
	}

	r = s % 11
	if (r == 0) || (r == 1) {
		fourteen = '0'
	} else {
		fourteen = rune((11 - r) + 48)
	}

	if (thirteen != rune(cnpj[12])) || (fourteen != rune(cnpj[13])) {
		ok = false
	}
	return
}

func ValidateCp(cpfcnpj string) (ok bool) {
	ok = true

	if _, err := strconv.Atoi(cpfcnpj); err != nil {
		ok = false
	}

	switch len(cpfcnpj) {
	case 11:
		ok = validateCpf(cpfcnpj)
	case 14:
		ok = validateCnpj(cpfcnpj)
	default:
		ok = false
	}

	return
}

func BuildReturnMessage(code int) (rtn []byte) {
	msg := struct {
		Codigo    int    `json:"codigo"`
		Mensagem  string `json:"mensagem"`
		Descricao string `json:"descricao"`
	}{
		code,
		MESSAGES[code],
		DESCRIPTIONS[code],
	}

	data := struct {
		Erro interface{} `json:"erro"`
	}{
		msg,
	}

	rtn, _ = json.Marshal(data)

	return
}

func IntSliceContains(sl []int, nm int) bool {
	for _, a := range sl {
		if a == nm {
			return true
		}
	}
	return false
}

func ValidateDoc(dbt map[string]interface{}) (ok bool, err error) {
	d, ok1 := dbt["documento"]
	e, ok2 := dbt["empid"]
	c, ok3 := dbt["clid"]
	v, ok4 := dbt["vencto"]

	ok5 := ok1 && ok2
	ok6 := ok3 && ok4
	ok7 := ok5 && ok6

	if !ok7 {
		ok = false
		err = errors.New("Erros nos parâmetros")
		return
	}

	doc := d.(string)

	_, err1 := strconv.ParseInt(doc, 10, 64)

	if err1 != nil {
		ok = false
		err = errors.New("Documento malformado")
		return
	}

	emp := e.(int)

	b1 := []byte{doc[0], doc[1]}
	i1, err2 := strconv.Atoi(string(b1[:]))

	if emp != i1 || err2 != nil {
		ok = false
		err = errors.New("Id da Empresa errado no documento")
		return
	}

	cli := c.(string)

	s1 := doc[2 : len(cli)+2]

	if s1 != cli {
		ok = false
		err = errors.New("Id do Cliente errado no documento")
		return
	}

	vct := v.(time.Time)

	i2, _ := strconv.Atoi(doc[len(cli)+2 : len(cli)+4])
	i3, _ := strconv.Atoi(doc[len(cli)+4 : len(cli)+6])

	if (i2 != vct.Day()) || (i3 != int(vct.Month())) {
		ok = false
		err = errors.New("Vencimento errado no documento")
		return
	}

	ok = true
	err = nil
	return
}

func ValidateDebt(debt map[string]interface{}) (ok bool, err error) {
	do, ok1 := debt["documento"]
	em, ok2 := debt["empid"]
	cl, ok3 := debt["clid"]
	no, ok4 := debt["nome"]
	te, ok5 := debt["telefone"]
	va, ok6 := debt["valor"]
	en, ok7 := debt["endereco"]
	ci, ok8 := debt["cidade"]
	ve, ok9 := debt["vencto"]
	cp, okA := debt["cpf_cnpj"]
	ca, okB := debt["cartao"]

	okC := ok1 && ok2 && ok3 && ok4
	okD := ok5 && ok6 && ok7 && ok8
	okE := ok9 && okA && okB && okC && okD

	if !okE {
		ok = false
		err = errors.New("Campos perdidos na fatura")
		return
	}

	var doc string
	switch do.(type) {
	case string:
		doc = do.(string)
	default:
		ok = false
		err = errors.New("Númerro de Documento de tipo inválido")
		return
	}
	if s := strings.TrimSpace(strings.Replace(doc, " ", "", -1)); s == "" {
		ok = false
		err = errors.New("Númerro de Documento vazio")
		return
	}

	var emp int
	switch em.(type) {
	case float64:
		emp = int(em.(float64))
	default:
		ok = false
		err = errors.New("Id da Empresa de tipo inválido")
		return
	}
	if emp <= 0 {
		ok = false
		err = errors.New("Id da Empresa de valor inválido")
		return
	}

	var cli int
	switch cl.(type) {
	case float64:
		cli = int(cl.(float64))
	default:
		ok = false
		err = errors.New("Id do Cliente de tipo inválido")
		return
	}
	if cli <= 0 {
		ok = false
		err = errors.New("Id do Cliente de valor inválido")
		return
	}

	var tel string
	switch te.(type) {
	case string:
		tel = te.(string)
	default:
		ok = false
		err = errors.New("Telefone do Cliente de tipo inválido")
		return
	}
	if s := strings.TrimSpace(strings.Replace(tel, " ", "", -1)); s == "" {
		ok = false
		err = errors.New("Telefone do Cliente vazio")
		return
	}
	if !ValidateTel(tel) {
		ok = false
		err = errors.New("Telefone do Cliente inválido")
		return
	}

	var nom string
	switch no.(type) {
	case string:
		nom = no.(string)
	default:
		ok = false
		err = errors.New("Nome do Cliente de tipo inválido")
		return
	}
	if s := strings.TrimSpace(strings.Replace(nom, " ", "", -1)); s == "" {
		ok = false
		err = errors.New("Nome do Cliente vazio")
		return
	}

	var val float32
	switch va.(type) {
	case float64:
		val = float32(cl.(float64))
	default:
		ok = false
		err = errors.New("Valor da fatura de tipo inválido")
		return
	}
	if val <= 0 {
		ok = false
		err = errors.New("Fatura com valor inválido")
		return
	}

	var car string
	switch ca.(type) {
	case string:
		car = ca.(string)
	default:
		ok = false
		err = errors.New("Cartão do Cliente de tipo inválido")
		return
	}
	if s := strings.TrimSpace(strings.Replace(car, " ", "", -1)); s == "" {
		ok = false
		err = errors.New("Cartão do Cliente vazio")
		return
	}
	if !ValidateCard(car) {
		ok = false
		err = errors.New("Cartão do Cliente inválido")
		return
	}

	var cpf_cnpj string
	switch cp.(type) {
	case string:
		cpf_cnpj = cp.(string)
	default:
		ok = false
		err = errors.New("CPF/CNPJ do Cliente de tipo inválido")
		return
	}
	if s := strings.TrimSpace(strings.Replace(cpf_cnpj, " ", "", -1)); s == "" {
		ok = false
		err = errors.New("CPF/CNPJ do Cliente vazio")
		return
	}
	if !ValidateCp(cpf_cnpj) {
		ok = false
		err = errors.New("CPF/CNPJ do Cliente inválido")
		return
	}

	var end string
	switch en.(type) {
	case string:
		end = en.(string)
	default:
		ok = false
		err = errors.New("Endereço do Cliente de tipo inválido")
		return
	}
	if s := strings.TrimSpace(strings.Replace(end, " ", "", -1)); s == "" {
		ok = false
		err = errors.New("Endereço do Cliente vazio")
		return
	}

	var cid string
	switch ci.(type) {
	case string:
		cid = ci.(string)
	default:
		ok = false
		err = errors.New("Cidade do Cliente de tipo inválido")
		return
	}
	if s := strings.TrimSpace(strings.Replace(cid, " ", "", -1)); s == "" {
		ok = false
		err = errors.New("Cidade do Cliente vazio")
		return
	}

	var s1 string
	switch ve.(type) {
	case string:
		s1 = ve.(string)
	default:
		ok = false
		err = errors.New("Vencimento da fatura de tipo inválido")
		return
	}
	if s := strings.TrimSpace(strings.Replace(s1, " ", "", -1)); s == "" {
		ok = false
		err = errors.New("Vencimento da fatura vazio")
		return
	}
	var ven time.Time
	var er error
	ven, er = time.Parse("02/01/2006 15:04:05", s1)
	if er != nil {
		ok = false
		err = errors.New("Vencimento com valor inválido")
	}

	dbt := make(map[string]interface{})
	dbt["documento"] = doc
	dbt["empid"] = emp
	dbt["clid"] = strconv.Itoa(cli)
	dbt["vencto"] = ven

	ok, err = ValidateDebt(dbt)

	return
}

func init() {
	lote = 0
}
