package main

import (
	"fmt"
	"github.com/cevaris/stockfighter"
	"os"
	"text/template"
	"time"
)

func main() {
	config := stockfighter.InitConfig(".env.yml",