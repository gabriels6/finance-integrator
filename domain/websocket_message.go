package domain

type WebSocketMessage struct {
	Action string          `json:"action"`
	Params WebSocketParams `json:"params"`
}

type SymbolParam struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange,omitempty"`
	MicCode  string `json:"mic_code,omitempty"`
	Type     string `json:"type,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type WebSocketParams struct {
	Symbols []SymbolParam `json:"symbols"`
}

type SuccessDetail struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	Country  string `json:"country"`
	Type     string `json:"type"`
}

type PriceEvent struct {
	Event     string          `json:"event"`
	Symbol    string          `json:"symbol"`
	Currency  string          `json:"currency"`
	Exchange  string          `json:"exchange"`
	Type      string          `json:"type"`
	Timestamp int64           `json:"timestamp"`
	Price     float64         `json:"price"`
	DayVolume int64           `json:"day_volume"`
	Success   []SuccessDetail `json:"success"`
	Fails     []SuccessDetail `json:"fails"`
}
