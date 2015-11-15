package utils

import (
	"encoding/json"
	"strconv"
	"strings"
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

func BuildReturnMessage(code int) (rtn string) {
	msg := struct {
		Codigo    int    `json:"codigo"`
		Mensagem  string `json:"mensagem"`
		Descricao string `json:"descricao"`
	}{
		code,
		MESSAGES[code],
		DESCRIPTIONS[code],
	}

	data, _ := json.Marshal(msg)
	rtn = string(data[:])

	return
}

func init() {
	lote = 0
}
