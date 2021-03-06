package model

type ReqExchangeKline struct {
	ExchangeId   int64  `json:"exchangeId" example:"2358885120275906"`
	Size         int32  `json:"size" example:"150"`
	Period       int32  `json:"period" example:"0"`
	CurrencyPair string `json:"currencyPair" example:"BTC_USDT"`
}
