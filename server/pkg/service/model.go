package service

const (
	_cotacaoUSDBRLURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

type CotacaoAtual struct {
	CotacaoDolarReal USDBRL `json:"USDBRL"`
}

type USDBRL struct {
	Code   string `json:"code,omitempty"`
	CodeIn string `json:"codein,omitempty"`
	Name   string `json:"name,omitempty"`
	High   string `json:"high,omitempty"`
	Low    string `json:"low,omitempty"`
	Bid    string `json:"bid,omitempty"`
}
