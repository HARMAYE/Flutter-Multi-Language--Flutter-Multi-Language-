package main

import (
	"time"
	"fmt"
	"github.com/cevaris/stockfighter"
)

func main() {
	fmt.Println(time.Now().Unix())

	config := stockfighter.InitConfig(".env.yml")

	fmt.Printf("config: %#v\n", config.ApiKey)

	api := stockfighter.InitApi(config, "HAE23155229")

	if value, err := api.HeartBeat(); err == nil {
		fmt.Printf("request: %#v\n", value)
	} else {
		fmt.Println(err)
	}

	if value, err := api.VenueHeartBeat("TESTEX"); err == nil {
		fmt.Printf("request: %#v\n", value)
	} else {
		fmt.Println(err)
	}

	if value, err := api.VenueStocks("TESTEX"); err == nil {
		fmt.Printf("request: %#v\n", value)
		fmt.Printf("%#v\n", value.Symbols[0].Name)
	} else {
		fmt.Println(err)
	}

	if value, err := api.StockOrderBook("TESTEX", "FOOBAR"); err == nil {
		fmt.Printf("request: %#v\n", value)
	} else {
		fmt.Println(err)
	}

	soReq := &stockfighter.StockOrderRequest{
		Accou