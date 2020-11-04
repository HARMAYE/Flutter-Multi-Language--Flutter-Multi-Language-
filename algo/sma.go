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
	o.data.PushBack(x)
}

func (o *SimpleMovingAvg) Value() int {
	var sum int = 0
	var currPeriod = 0
	for e := o.data.Front(); e != nil && currPeriod <= o.period; e = e.Next() {
		sum += e.Value.(int)
		currPeriod += 1
	}
	return sum / o.period
}

func (o *SimpleMovingAvg) String() string {
	return fmt.Sprintf(
		"SimpleMovingAvg(Period:%d Value:%d Trend:%d Slice:%+v)",
		o.period, o.Value(), o.Trend(), o.slice(),
	)
}

func (o *Sim