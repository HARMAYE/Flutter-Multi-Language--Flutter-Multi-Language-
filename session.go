
package stockfighter

import (
	"sync"
	"fmt"
	"encoding/json"
)

type Session struct {
	Cash          int
	NAV           int
	Position      int
	TotalPosition int
	Venue         string
	api           *Api
	config        *config
	mutex         *sync.RWMutex
	LatestQuote   *StockQuote
	quoteChan     chan *StockQuote
	fillChan      chan *Execution
}