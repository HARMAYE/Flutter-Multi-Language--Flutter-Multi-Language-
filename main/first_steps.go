package main

import (
	"fmt"
	"github.com/cevaris/stockfighter"
	"time"
)

var (
	Account = "MS48925855"
	Venue = "WUTMEX"
	Symbol = "EYNI"
)

var (
	config = stockfighter.InitConfig(".env.yml", Account)
	