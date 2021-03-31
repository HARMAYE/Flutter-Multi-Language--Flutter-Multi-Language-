
package algo

import (
	"testing"
	"reflect"
	"github.com/cevaris/stockfighter"
)

var testPeriod int = 5

func TestSimpleMovingAvg(t *testing.T) {
	sma := InitSimpleMovingAvg(testPeriod)
	if sma.Value() != 0 {
		t.Error("bad value", sma.Value())
	}
}

func TestSimpleMovingAvg_Push(t *testing.T) {
	sma := InitSimpleMovingAvg(testPeriod)
	for i := 1; i <= testPeriod; i++ {
		sma.Push(i)