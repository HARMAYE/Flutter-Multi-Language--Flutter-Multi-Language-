package stockfighter

import (
	"github.com/spf13/viper"
	"bytes"
	"io/ioutil"
	"fmt"
)

type config struct {
	ApiKey  string
	Account string
}

func InitConfig(filepath string, account string) *config {
	viper.SetConfigType("yaml")

	if bs, err := ioutil.ReadFile(filepath); err == nil {
		viper.ReadConfig