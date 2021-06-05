
package main

import (
	"fmt"
	"github.com/cevaris/stockfighter"
	"log"
	"time"
	"math/rand"
)

var (
	venue = "TJMEX"
	symbol = "MBL"
	account = "KEG25461931"

	config = stockfighter.InitConfig(".env.yml", account)
	api = stockfighter.InitApi(config)
	random = rand.New(rand.NewSource(time.Now().Unix()))
)

func main() {
	var currShares = 25800
	targetShares := 100000
	maxShareStep := 30000
	maxAskStep := 100

	api.IsExchangeHealthy()
	baseQuote := api.GetBaseQuote()

	for currShares < targetShares {
		shareStep := random.Intn(maxShareStep)
		askStep := random.Intn(maxAskStep)

		currQuote, err := api.StockQuote(venue, symbol)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Curr Ask", currQuote.Ask)

		if currQuote.Ask > 0 && currQuote.Ask < 5100 && (currQuote.Ask - askStep) > 0 {
			soReq := &stockfighter.StockOrderRequest{