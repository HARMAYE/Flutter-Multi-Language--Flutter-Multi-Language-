
package stockfighter

import (
	"testing"
)

var testVenue = "TestVenue"

func TestInitSession(t *testing.T) {
	session := InitSession(&config{}, testVenue)
	if session.Position != 0 {
		t.Error("invalid position", session.Position)
	}
	if session.Cash != 0 {
		t.Error("invalid cash", session.Cash)
	}
	if session.NAV != 0 {
		t.Error("invalid nav", session.NAV)
	}
}

func TestSessionUpdate(t *testing.T) {
	session := InitSession(&config{}, testVenue)
	session.LatestQuote = &StockQuote{Last: 100}
	session.Update(&StockOrderAccountStatus{
		Ok: true,
		Venue: testVenue,
		Orders: []*StockOrder{
			&StockOrder{
				Ok: true,
				Direction: DirectionBuy,
				Fills: []*Fill{