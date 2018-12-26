package main

import (
	"flag"
	"fmt"
	"yicfu.com/tax/y2019/lib"
)

var (
	monthlyIncome = flag.Float64("mi", 10000, "月收入（税前）")
	taxBase       = flag.Float64("tb", 5000, "个税起征点")
	insurrance    = flag.Float64("xj", 500, "五险一金")
	months        = flag.Int("m", 12, "计算多少个月")

	// TODO(fuyc): support different income and tax base for various months.
)

func main() {
	flag.Parse()
	result, err := lib.Calc(lib.Request{
		MonthlyIncome: *monthlyIncome,
		TaxBase:       *taxBase,
		Insurrance:    *insurrance,
		Months:        *months,
	})
	fmt.Println(result, err)
}
