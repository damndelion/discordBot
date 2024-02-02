package dto

// CoinGeckoResponse dto object to get price of BTC
type CoinGeckoResponse struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}
