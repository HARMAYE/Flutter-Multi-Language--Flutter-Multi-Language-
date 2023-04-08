
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
	}
}

func (o *Session) String() string {
	j, _ := json.Marshal(o)
	return fmt.Sprintf("Session(%s)", string(j))
}

func (o *Session) Update(status *StockOrderAccountStatus) {
	var sumPositionSecured int = 0
	var sumPositionNonSecured int = 0
	var sumCash int = 0

	if !status.Ok {
		return
	}

	for _, so := range status.Orders {
		if !so.Ok {
			continue
		}