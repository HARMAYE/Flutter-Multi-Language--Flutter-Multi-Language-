
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

func (s *Api) HeartBeat() (*HeartBeat, error) {
	buffer, err := s.GetRequest("ob/api/heartbeat")
	if err != nil {
		return nil, err
	}

	var value *HeartBeat
	jsonErr := json.Unmarshal(buffer, &value)

	if jsonErr == nil {
		return value, nil
	}

	return nil, jsonErr
}

func (s *Api) VenueHeartBeat(venue string) (*VenueHeartBeat, error) {
	buffer, err := s.GetRequest(fmt.Sprintf("ob/api/venues/%s/heartbeat", venue))
	if err != nil {
		return nil, err
	}

	var value *VenueHeartBeat
	jsonErr := json.Unmarshal(buffer, &value)
	if jsonErr == nil {
		return value, nil
	}

	return nil, jsonErr
}

func (s *Api) VenueStocks(venue string) (*VenueStocks, error) {
	buffer, err := s.GetRequest(fmt.Sprintf("ob/api/venues/%s/stocks", venue))
	if err != nil {
		return nil, err
	}

	var value *VenueStocks

	jsonErr := json.Unmarshal(buffer, &value)
	if jsonErr == nil {
		return value, nil
	}

	return nil, jsonErr
}

func (s *Api) StockOrderBook(venue string, stock string) (*StockOrderBook, error) {
	buffer, err := s.GetRequest(fmt.Sprintf("ob/api/venues/%s/stocks/%s", venue, stock))

	if err != nil {
		return nil, err
	}

	var value *StockOrderBook

	jsonErr := json.Unmarshal(buffer, &value)
	if jsonErr == nil {
		return value, nil
	}

	return nil, jsonErr
}

func (s *Api) StockOrder(soReq *StockOrderRequest) (*StockOrder, error) {
	url := fmt.Sprintf(
		"ob/api/venues/%s/stocks/%s/orders", soReq.Venue, soReq.Stock,
	)
	buffer, err := s.PostRequest(url, soReq)
	if err != nil {
		return nil, err