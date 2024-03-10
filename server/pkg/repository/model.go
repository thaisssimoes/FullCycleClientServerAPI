package repository

const (
	_cotacaoUSDBRLURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

type CotacaoAtual struct {
	CotacaoDolarReal USDBRL `json:"USDBRL"`
}

type USDBRL struct {
	Code       string `json:"code,omitempty" db:"code"`
	CodeIn     string `json:"codein,omitempty" db:"code_in"`
	Name       string `json:"name,omitempty" db:"name"`
	High       string `json:"high,omitempty" db:"high"`
	Low        string `json:"low,omitempty" db:"low"`
	Bid        string `json:"bid,omitempty" db:"bid"`
	VarBid     string `json:"varBid" db:"var_bid"`
	PctChange  string `json:"pctChange" db:"pct_change"`
	Ask        string `json:"ask" db:"ask"`
	Timestamp  string `json:"timestamp" db:"timestamp"`
	CreateDate string `json:"create_date" db:"create_date"`
}
