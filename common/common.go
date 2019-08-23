package common

// Sub 结构
type Sub struct {
	Sub string `json:"sub"`
	ID  string `json:"id"`
}

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

// DeleteByName xxs
func DeleteByName(subs []interface{}, name string) []interface{} {
	for i := 0; i < len(subs); i++ {
		if subs[i].(map[string]string)["Sub"] == name {
			subs = append(subs[:i], subs[i+1:]...)
		}
	}
	return subs
}
