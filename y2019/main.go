package main

import (
	"flag"
	"fmt"
)

var (
	monthlyIncome = flag.Float64("mi", 10000, "月收入（税前）")
	taxBase       = flag.Float64("tb", 5000, "个税起征点")
	insurrance    = flag.Float64("xj", 500, "五险一金")
	months        = flag.Float64("m", 12, "计算多少个月")
)

type Deduct struct {
	Level           int
	TaxedIncomeTop  float64
	DeductRate      float64
	QuickCalcDeduct float64
}

type Result struct {
	// 1. 逐月的扣税
	// 2. 逐月的到手数
	// 3. 合计的到手数
	Deducts []float64
	Obtains []float64
	Total   float64
}

var (
	deducts = []Deduct{
		Deduct{1, 36000, 3, 0},
		Deduct{2, 144000, 10, 2520},
		Deduct{3, 300000, 20, 16920},
		Deduct{4, 420000, 25, 31920},
		Deduct{5, 660000, 30, 52920},
		Deduct{6, 960000, 35, 85920},
		Deduct{7, 9999990000, 45, 181920},
	}
)

func main() {
	flag.Parse()
	result, err := calc()
	fmt.Println(result, err)
}

func calc() (Result, error) {
	//TODO
	res:= Result{}
	return res , nil
}
