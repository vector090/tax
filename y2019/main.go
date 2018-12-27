package main

import (
	"flag"
	"fmt"
	"yicfu.com/tax/y2019/lib"
)

var (
	monthlyIncome     = flag.Float64("mi", 10000, "月收入（税前）")
	taxBase           = flag.Float64("tb", 5000, "个税起征点")
	insurrance        = flag.Float64("xj", 500, "个税扣除（五险一金）")
	additionalTaxFree = flag.Float64("atf", 0, "个税专项附加扣除")
	months            = flag.Int("m", 12, "计算多少个月")

	// TODO(fuyc): support different income and tax base for various months.
)

func main() {
	flag.Parse()
	result, err := lib.Calc(lib.Request{
		MonthlyIncome:     *monthlyIncome,
		TaxBase:           *taxBase,
		Insurrance:        *insurrance,
		AdditionalTaxFree: *additionalTaxFree,
		Months:            *months,
	})
	if nil != err {
		panic(err)
	}
	//fmt.Printf("%+v\n", result)
	fmt.Println("Months\t", *months)
	//fmt.Println("Deducts\t", result.Deducts)
	//fmt.Println("Obtains\t", result.Obtains)
	for m := 0; m < *months; m++ {
		fmt.Printf("%d月\t", m+1)
		fmt.Printf("当月扣个税 %.2f\t", result.Deducts[m])
		fmt.Printf("当月到手 %.2f\t", result.Obtains[m])
		fmt.Println()
	}
	fmt.Println("总到手\t", result.TotalObtains)
}
