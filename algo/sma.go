package algo

import (
	"container/list"
	"sync"
	"github.com/cevaris/stockfighter"
	"fmt"
)

/*
http://www.investopedia.com/articles/active-trading/012815/top-technical-indicators-scalping-trading-strategy.asp
http://www.investopedia.com/articles/active-trading/010116/perfect-moving-averages-day-trading.asp
http://www.investopedia.com/articles/active-trading/012815/top-technical-indicators-scalping-trading-strategy.asp
http://www.day-trading-stocks.org/moving-average-crossover.html
 */

type SimpleMovingAvg struct {
	data   list.List
	period int
}

func InitSimpleMovingAvg(period int) *SimpleMovingAvg {
	return &SimpleMovingAvg{period: period}
}

func (o *SimpleMovingAvg) Push(x int) {
	if o.data.Len() >= o.period {
		o.data.Remove(o.data.Front())
	}
	o.data.Pus