
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