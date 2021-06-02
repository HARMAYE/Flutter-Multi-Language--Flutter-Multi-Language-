
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
