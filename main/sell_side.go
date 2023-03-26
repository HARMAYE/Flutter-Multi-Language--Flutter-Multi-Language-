
package main

import (
	"fmt"
	"github.com/cevaris/stockfighter"
	"github.com/cevaris/stockfighter/algo"
	"time"
)

var (
	Account = "ATE44401158"
	Venue = "LVBTEX"
	Symbol = "RPM"
)

var (
	config = stockfighter.InitConfig(".env.yml", Account)
	api = stockfighter.InitApi(config)
	session = stockfighter.InitSession(config, Venue)
	smaTri = algo.InitSmaTriple(5, 10, 13)
)

func main() {
	api.IsExchangeHealthy()
	session.Observe(Symbol)

	go positionWorker()

	for {
		fmt.Println(session)
		time.Sleep(3 * time.Second)
	}
}

func positionWorker() {
	for {
		quote := session.LatestQuote

		if quote != nil {
			smaTri.Push(quote.Last)
