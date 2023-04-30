
package stockfighter

import (
	"time"
	"fmt"
	"encoding/json"
)

const (
	DirectionBuy string = "buy"
	DirectionSell string = "sell"

	OrderLimit string = "limit"
	OrderMarket string = "market"
	OrderFillOrKill string = "fill-or-kill"
	OrderImmediateOrCancel string = "immediate-or-cancel"

	SignalBuy int = 1
	SignalUnknown int = 0
	SignalSell int = -1

	TrendUp int = 1
	TrendUnknown int = 0
	TrendDown int = -1
)

type Symbol struct {
	Name   string
	Symbol string
}

type Bid struct {
	Price int
	Qty   int
	IsBuy bool
}

type Ask struct {