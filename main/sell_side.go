
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