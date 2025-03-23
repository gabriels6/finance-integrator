package apis

import "fmt"

func GetBcbExchangeRate(currency string, date string) []byte {
	return CallApi(fmt.Sprintf("https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/odata/CotacaoMoedaDia(moeda=@moeda,dataCotacao=@dataCotacao)?@moeda='%s'&@dataCotacao='%s'&$top=1&$skip=0&$format=json&$select=cotacaoCompra,cotacaoVenda,dataHoraCotacao", currency, date))
}

func GetBcbExchangeRateByPeriod(currency string, startDate string, endDate string) []byte {
	return CallApi(fmt.Sprintf("https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/odata/CotacaoMoedaPeriodo(moeda=@moeda,dataInicial=@dataInicial,dataFinalCotacao=@dataFinalCotacao)?@moeda='%s'&@dataInicial='%s'&@dataFinalCotacao='%s'&$top=1100&$format=json&$select=cotacaoCompra,cotacaoVenda,dataHoraCotacao", currency, startDate, endDate))
}
