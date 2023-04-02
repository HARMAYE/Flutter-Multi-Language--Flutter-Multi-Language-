
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

func InitSession(config *config, venue string) *Session {
	return &Session{
		NAV: 0,
		Position: 0,
		TotalPosition: 0,
		Cash: 0,
		Venue: venue,
		api: InitApi(config),
		config: config,
		mutex: &sync.RWMutex{},
		quoteChan: make(chan *StockQuote, 100),
		fillChan: make(chan *Execution, 100),