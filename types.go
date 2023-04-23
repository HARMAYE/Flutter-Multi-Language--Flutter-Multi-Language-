
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