package apis

import "fmt"

func GetBcbExchangeRate(currency string, date string) []byte {
	return CallApi(fmt.Sprintf("https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/odata/CotacaoMoedaDia(moeda=@moeda,dataCotacao=@dataCotacao)?@moeda='%s'&@dataCotacao='%s'&$top=1&$skip=0&$format=json&$select=cotacaoCompra,cotacaoVenda,dataHoraCotacao", currency, date))
}
