package main

import (
	"fmt"
	"github.com/cevaris/stockfighter"
	"os"
	"text/template"
	"time"
)

func main() {
	config := stockfighter.InitConfig(".env.yml", "SDS22054882")
	fmt.Printf("config: %#v\n", config.ApiKey)

	api := stockfighter.InitA