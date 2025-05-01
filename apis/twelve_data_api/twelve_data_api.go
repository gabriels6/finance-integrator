package twelvedataapi

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gabriels6/finance-integrator/apis"
	"github.com/gabriels6/finance-integrator/domain"
	"github.com/gabriels6/finance-integrator/utils"
	"github.com/gabriels6/finance-integrator/websocket"
)

const TWELVE_DATA_BASE_URL = "https://api.twelvedata.com"

const TWELVE_DATA_WS_URL = "wss://ws.twelvedata.com/v1/quotes/price?apikey="

var wsPricesData map[string]domain.PriceEvent = make(map[string]domain.PriceEvent)

var wsClient *websocket.WebSocketClient

func GetSeries(symbols string) []byte {
	return apis.CallApi(fmt.Sprintf("%s/time_series?symbol=%s&apikey=%s&interval=1day", TWELVE_DATA_BASE_URL, symbols, utils.GetEnv("TWELVE_DATA_API_KEY")))
}

func GetEodPrices(symbols string) []byte {
	// return []byte(fmt.Sprintf("{\"req_1\":{ \"/eod?symbol=%s&apikey=%s\"} }", symbols, utils.GetEnv("TWELVE_DATA_API_KEY")))
	return apis.PostApi(fmt.Sprintf("%s/batch", TWELVE_DATA_BASE_URL), fmt.Sprintf("{\"req_1\": { \"url\": \"/eod?symbol=%s&apikey=%s\"} }", symbols, utils.GetEnv("TWELVE_DATA_API_KEY")))
}

func SearchSymbol(keyword string) []byte {
	return apis.CallApi(fmt.Sprintf("%s/symbol_search?symbol=%s&apikey=%s", TWELVE_DATA_BASE_URL, keyword, utils.GetEnv("TWELVE_DATA_API_KEY")))
}

func GetStocks() []byte {
	return apis.CallApi(fmt.Sprintf("%s/stocks?apikey=%s", TWELVE_DATA_BASE_URL, utils.GetEnv("TWELVE_DATA_API_KEY")))
}

func GetStock(symbol string) []byte {
	return apis.CallApi(fmt.Sprintf("%s/stocks?symbol=%s&apikey=%s", TWELVE_DATA_BASE_URL, symbol, utils.GetEnv("TWELVE_DATA_API_KEY")))
}

func GetETFs() []byte {
	return apis.CallApi(fmt.Sprintf("%s/etfs?apikey=%s", TWELVE_DATA_BASE_URL, utils.GetEnv("TWELVE_DATA_API_KEY")))
}

func GetWsRealtimePrices() []byte {
	pricesData, err := json.Marshal(wsPricesData)
	if err != nil {
		return []byte(fmt.Sprintf(`{ "error": "error marshalling prices: %v" }`, pricesData))
	}
	return pricesData
}

func GatherWebsocketRealTimeQuotes(symbols string) error {
	if wsClient != nil {
		wsClient.Close()
	}
	pricesClient, err := websocket.NewWebSocketClient(fmt.Sprintf("%s%s", TWELVE_DATA_WS_URL, utils.GetEnv("TWELVE_DATA_API_KEY")))
	if err != nil {
		return fmt.Errorf("error creating ws client: %v", err)
	}
	wsClient = pricesClient
	symbolsData := []domain.SymbolParam{}
	for _, symbol := range strings.Split(symbols, ",") {
		symbolsData = append(symbolsData, domain.SymbolParam{
			Symbol:   symbol,
			Exchange: "NASDAQ",
		})
	}
	message := domain.WebSocketMessage{
		Action: "subscribe",
		Params: domain.WebSocketParams{
			Symbols: symbolsData,
		},
	}
	stringMessage, err := json.Marshal(message)
	fmt.Print("Subscribe message: %s", string(stringMessage))
	if err != nil {
		return fmt.Errorf("error marshalling subscribe message: %v", err)
	}
	pricesClient.Subscribe(string(stringMessage))
	go func() {
		defer pricesClient.Close()
		pricesClient.ReceiveMessages(HandleQuoteMessage)
	}()
	return nil
}

func HandleQuoteMessage(messageType int, message []byte) {
	fmt.Printf("Event received: %s\n", string(message))
	event, err := ParseQuoteMessage(message)
	if err != nil {
		fmt.Printf("Error parsing quote message: %v\n", err)
		return
	}
	switch strings.ToLower(event.Event) {
	case "subscribe-status":
		var sucessSymbols, failedSymbols string
		for _, detail := range event.Success {
			sucessSymbols = sucessSymbols + "," + detail.Symbol
		}
		for _, detail := range event.Fails {
			failedSymbols = failedSymbols + "," + detail.Symbol
		}
		fmt.Printf("Subscribe status: sucess: %v; Failed: %v", sucessSymbols, failedSymbols)
		break
	case "price":
		wsPricesData[event.Symbol] = event
		break
	}
}

func ParseQuoteMessage(message []byte) (domain.PriceEvent, error) {
	var prices domain.PriceEvent
	err := json.Unmarshal(message, &prices)
	if err != nil {
		return domain.PriceEvent{}, fmt.Errorf("error unmarshalling price event: %v", err)
	}
	return prices, nil
}
