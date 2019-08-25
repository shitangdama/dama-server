package common

// Sub 结构
type Sub struct {
	Sub string `json:"sub"`
	ID  string `json:"id"`
}

// KTick 结构
type KTick struct {
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Vol    float64 `json:"vol"`
	Count  int     `json:"count"`
	Amount float64 `json:"amount"`
}

// KTicker 结构
type KTicker struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick KTick  `json:"tick"`
}

// DTick xx
type DTick struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

// DTicker 结构
type DTicker struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick DTick  `json:"tick"`
}

// DeTick 结构
type DeTick struct {
	ID int `json:"id"`
	Ts int `json:"ts"`

	Count  int     `json:"count"`
	Amount float64 `json:"amount"`

	Open float64 `json:"open"`
	Low  float64 `json:"low"`
	High float64 `json:"high"`
	Vol  float64 `json:"vol"`
}

// DeTicker Detail
type DeTicker struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick DeTick `json:"tick"`
}

// TTicker trade
type TTicker struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick TTick  `json:"tick"`
}

// TTick xx
type TTick struct {
	ID   int     `json:"id"`
	Ts   int     `json:"ts"`
	Data []Trade `json:"data"`
}

// Trade xx
type Trade struct {
	Amount    float64 `json:"amount"`
	Ts        int     `json:"ts"`
	ID        int     `json:"id"`
	Price     float64 `json:"price"`
	Direction string  `json:"direction"`
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
