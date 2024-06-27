package models

type Message struct {
	Content string
}

type Ticker struct {
	Symbol           string `json:"s"`
	Change           string `json:"p"`
	ChangePct        string `json:"P"`
	WeightedAvgPrice string `json:"w"`
	PrevClosePrice   string `json:"x"`
	LastPrice        string `json:"c"`
	LastQty          string `json:"Q"`
	BidPrice         string `json:"b"`
	BidQty           string `json:"B"`
	AskPrice         string `json:"a"`
	AskQty           string `json:"A"`
	OpenPrice        string `json:"o"`
	HighPrice        string `json:"h"`
	Volume           string `json:"v"`
	QuoteVolume      string `json:"q"`
	OpenTime         int64  `json:"O"`
	CloseTime        int64  `json:"C"`
	TotalTrades      int64  `json:"n"`
}

type ProcessedTicker struct {
	Symbol    string
	Name      string
	Price     float64
	Change    float64
	ChangePct float64
	Volume    float64
}
