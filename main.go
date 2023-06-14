package main

import (
	"fmt"
	"time"

	"github.com/walbety/go-fii/internal/canonical"
	"github.com/walbety/go-fii/internal/integration/infomoney"
)

var (
	infomoneyService infomoney.Infomoney
)

func main() {
	var err error
	infomoneyService = infomoney.Service{}

	BBPO := canonical.FundoImob{Name: "BBPO", Ticker: "BBPO11"}

	MXRF := canonical.FundoImob{Name: "MXRF", Ticker: "MXRF11"}

	BBPO.Dados, err = infomoneyService.GetYieldTickersFromFII(BBPO.Ticker)
	if err != nil {
		fmt.Printf("error at retrieving %s", MXRF.Name)
	}
	MXRF.Dados, err = infomoneyService.GetYieldTickersFromFII(MXRF.Ticker)
	if err != nil {
		fmt.Printf("error at retrieving %s", MXRF.Name)
	}

	now := time.Now()
	oneMonthBefore := now.AddDate(0, 0, -90)

	BBPO.Valores, err = infomoneyService.GetValuesHistoric(BBPO.Ticker, oneMonthBefore, now)

	// fmt.Println(BBPO.Dados)
	// fmt.Println(MXRF.Dados)
	gerarRelatorioDiario(BBPO)

}

type Relacao struct {
	Nome  string
	Valor float32
	Desc  string
}

func gerarRelatorioDiario(fundo canonical.FundoImob) string {
	tickers := fundo.Valores.DataValor
	size := len(tickers)
	// j := 0

	mapaDeRelacoes := make(map[int]float32)

	iRel := 1
	for j := 1; j <= size; j++ {

		mapaDeRelacoes[iRel] = 1 - (tickers[size-j].Valor / tickers[size-1].Valor)
		iRel++
	}

	fmt.Printf("relação a 1 dia atrás: %v", mapaDeRelacoes[0])

	for i, relacao := range mapaDeRelacoes {
		fmt.Printf("relação entre hoje e %v dias: %v\n", i, relacao)
	}

	return ""
}
