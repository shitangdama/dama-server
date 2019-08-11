package common

// Tick 结构
type Tick struct {
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Vol    float64 `json:"vol"`
	Count  int     `json:"count"`
	Amount float64 `json:"amount"`
}

// Ticker 结构
type Ticker struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick Tick   `json:"tick"`
}
