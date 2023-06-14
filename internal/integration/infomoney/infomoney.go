package infomoney

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/walbety/go-fii/internal/canonical"
)

const (
	URL_INFOMONEY_HISTORICO = "https://fii-api.infomoney.com.br/api/v1/fii/provento/historico?Ticker="
	URL_INFOMONEY_VALORES   = "https://fii-api.infomoney.com.br/api/v1/fii/cotacao/historico/grafico"
)

type Infomoney interface {
	GetYieldTickersFromFII(Fundo string) ([]canonical.RendimentoTicker, error)
	GetValuesHistoric(name string, dataInicio time.Time, dataFim time.Time) (canonical.ValorTicker, error)
}

type Service struct {
}

func (s Service) GetYieldTickersFromFII(Fundo string) ([]canonical.RendimentoTicker, error) {

	url := fmt.Sprint(URL_INFOMONEY_HISTORICO, Fundo)

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("error at Get")
		return []canonical.RendimentoTicker{}, err
	}

	var rendimentos []canonical.RendimentoTicker

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error at readall")
		return []canonical.RendimentoTicker{}, err
	}

	err = json.Unmarshal(body, &rendimentos)
	if err != nil {
		fmt.Println("error at Unmarshal")
		return []canonical.RendimentoTicker{}, err
	}

	return rendimentos, nil
}

func (s Service) GetValuesHistoric(name string, dataInicio time.Time, dataFim time.Time) (canonical.ValorTicker, error) {
	dateFormat := "02-01-2006"

	baseURL, _ := url.Parse(URL_INFOMONEY_VALORES)
	params := url.Values{}
	params.Add("Ticker", name)
	params.Add("DataInicio", dataInicio.Format(dateFormat))
	params.Add("DataFim", dataFim.Format(dateFormat))
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		fmt.Printf("error at get")
		return canonical.ValorTicker{}, err
	}
	defer resp.Body.Close()

	var valor canonical.ValorTicker

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error at readall")
		return canonical.ValorTicker{}, err
	}

	err = json.Unmarshal(body, &valor)
	if err != nil {
		fmt.Println("error at Unmarshal VALOOR")
		return canonical.ValorTicker{}, err
	}

	return valor, nil
}
