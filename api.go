
package stockfighter

import (
	"github.com/franela/goreq"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"golang.org/x/net/websocket"
	"log"
	"io"
)

/*
Web Sockets: https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/08.2.html
*/

var (
	PRINT_JSON_RESPONSE = false
)

type Api struct {
	Config *config
}

func InitApi(config *config) *Api {
	return &Api{
		Config: config,
	}
}