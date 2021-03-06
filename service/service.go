package service

const (
	token = "token"
	//rootUrl = "http://bitesla-service-api:8090"
	rootUrl = "http://127.0.0.1:8090"
)

type Service struct {
	ApiKey        string
	ApiSecret     string
	CurrentUserId int64
	Token         string
	ExchangeId    int64
	TraderId      int64
	StrategyId    int64
	RunTimes      int64 //该策略运行了多少遍
}
