package canonical

type FundoImob struct {
	Name    string
	Ticker  string
	Dados   []RendimentoTicker
	Valores ValorTicker
}

type RendimentoTicker struct {
	Ticker     string  `json:"ticker"`
	Rendimento float32 `json:"rendimento"`
	Yield      float32 `json:"yield"`
	Data       string  `json:"data"`
}

type ValorTicker struct {
	Ticker    string `json:"ticker"`
	DataValor []DataValor
}

type DataValor struct {
	Data  string
	Valor float32 `json:"valor"`
}

/*
"ticker": "BBPO11",
    "dataValor": [
        {
            "data": "10-05-2023T00:00:00",
            "valor": 88.69
        },
        {
            "data": "11-05-2023T00:00:00",
            "valor": 88.7
        },
*/
