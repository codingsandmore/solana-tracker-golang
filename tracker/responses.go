package tracker

type RateResponse struct {
	AmountIn       float64 `json:"amountIn"`
	AmountOut      float64 `json:"amountOut"`
	MinAmountOut   float64 `json:"minAmountOut"`
	CurrentPrice   float64 `json:"currentPrice"`
	ExecutionPrice float64 `json:"executionPrice"`
	PriceImpact    float64 `json:"priceImpact"`
	Fee            float64 `json:"fee"`
	BaseCurrency   struct {
		Decimals int    `json:"decimals"`
		Mint     string `json:"mint"`
	} `json:"baseCurrency"`
	QuoteCurrency struct {
		Decimals int    `json:"decimals"`
		Mint     string `json:"mint"`
	} `json:"quoteCurrency"`
	PlatformFee      float64 `json:"platformFee"`
	PlatformFeeUI    float64 `json:"platformFeeUI"`
	IsJupiter        bool    `json:"isJupiter"`
	RawQuoteResponse struct {
		InputMint            string  `json:"inputMint"`
		InAmount             string  `json:"inAmount"`
		OutputMint           string  `json:"outputMint"`
		OutAmount            string  `json:"outAmount"`
		OtherAmountThreshold string  `json:"otherAmountThreshold"`
		SwapMode             string  `json:"swapMode"`
		SlippageBps          int     `json:"slippageBps"`
		PlatformFee          any     `json:"platformFee"`
		PriceImpactPct       float64 `json:"priceImpactPct,string"`
		RoutePlan            []struct {
			SwapInfo struct {
				AmmKey     string `json:"ammKey"`
				Label      string `json:"label"`
				InputMint  string `json:"inputMint"`
				OutputMint string `json:"outputMint"`
				InAmount   string `json:"inAmount"`
				OutAmount  string `json:"outAmount"`
				FeeAmount  string `json:"feeAmount"`
				FeeMint    string `json:"feeMint"`
			} `json:"swapInfo"`
			Percent int `json:"percent"`
		} `json:"routePlan"`
		ContextSlot int     `json:"contextSlot"`
		TimeTaken   float64 `json:"timeTaken"`
	} `json:"rawQuoteResponse"`
}

type SwapResponse struct {
	Txn  string       `json:"txn"`
	Rate RateResponse `json:"rate"`
}
