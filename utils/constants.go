package utils

const (
	OPERATION_NOT_ALLOWED = 1001
	RESOURCE_NOT_FOUND    = 1002
	HTTP_SERVER_ERROR     = 1003

	INVALIDE_PARAMETER_TYPE  = 2001
	INVALIDE_PARAMETER_VALUE = 2002

	TERMINAL_NOT_FOUND = 3001
	CLIENT_NOT_FOUND   = 3002
	DEBT_NOT_FOUND     = 3003

	NO_DEBTS_LEFT           = 4001
	DEBT_ALREADY_PAID       = 4002
	TERMINAL_ALREADY_OPENED = 4003
	TERMINAL_ALREADY_CLOSED = 4004
	TERMINAL_NOT_OPENED     = 4005
	TERMINAL_NOT_CLOSED     = 4006

	INVALIDE_OPENING_DATE   = 5001
	INVALIDE_PAYMENT_DATE   = 5002
	INVALIDE_PAID_AMOUNT    = 5003
	INVALIDE_TERMINAL_VALUE = 5004

	IMPOSSIBLE_TO_CONFIRM = 6001
	IMPOSSIBLE_TO_SEARCH  = 6002
)

var (
	MESSAGES = map[int]string{
		1001: "Operação Não Permitida",
		1002: "Recurso Não Encontrado",
		1003: "Erro no Servidor",
		2001: "Parâmetro de Tipo Inválido",
		2002: "Parâmetro de Valor Inválido",
		3001: "Terminal não encontrado",
		3002: "Cliente não Encontrado",
		3003: "Débito não Encontrado",
		4001: "Não há Débitos",
		4002: "Débito já Pago",
		4003: "Terminal já Aberto",
		4004: "Terminal já fechado",
		4005: "Terminal não Aberto",
		4006: "Terminal não fechado",
		5001: "Data de Abertura Inválida",
		5002: "Data de Pagamento Inválida",
		5003: "Valor Pago Inválido",
		5004: "Número de Terminal Inválido",
		6001: "Impossível de Confirmar",
		6002: "Impossível de Pesquisar",
	}

	DESCRIPTIONS = map[int]string{
		1001: "O servidor não é perimitido executar a operação que deseja acessar",
		1002: "O recurso que deseja acessar não foi encontrado",
		1003: "O servidor apresentou um erro, tente essa operação novamente mais tarde",
		2001: "Os parâmetros da requisição não são de um tipo válido",
		2002: "Os parâmetros da requisição não possuem valores válidos",
		3001: "O terminal com o qual deseja fazer a operação não foi encontrado",
		3002: "O cliente a qual busca se refere não foi encontrado",
		3003: "O débito em questão não foi encontrado",
		4001: "O cliente em questão não possui débitos em aberto",
		4002: "O débito em questão já foi pago",
		4003: "O terminal em questão já foi aberto",
		4004: "O terminal em questão já foi fechado",
		4005: "O terminal em questão não foi aberto",
		4006: "O terminal em questão não foi fechado",
		5001: "A data de abertura do caixa não possui um valor válido",
		5002: "A data de pagamento do débito não possui um valor válido",
		5003: "O valor pago pelo débito não possui um valor válido",
		5004: "O número de terminal enviado é inválido",
		6001: "Não foi possível confirmar o pagamento este devido a um erro no servidor",
		6002: "Não foi possível realizar esta pesquisa devido a um erro no servidor",
	}

	PRINTERS = map[int]string{
		1:   "P25_029363_01",
		111: "P25_032529_01",
		130: "P25_032510_01",
	}

	ADDRESSES = map[int]string{
		1:   "00:08:1B:95:73:DC",
		111: "00:08:1B:95:3B:18",
		130: "00:08:1B:95:39:E9",
	}

	INVALIDES = []string{
		"00000000000",
		"11111111111",
		"22222222222",
		"33333333333",
		"44444444444",
		"55555555555",
		"66666666666",
		"77777777777",
		"88888888888",
		"99999999999",
		"00000000000000",
		"11111111111111",
		"22222222222222",
		"33333333333333",
		"44444444444444",
		"55555555555555",
		"66666666666666",
		"77777777777777",
		"88888888888888",
		"99999999999999",
	}

	TERMS = []int{1, 111, 130}
)
