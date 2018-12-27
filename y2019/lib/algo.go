package lib

import (
	"errors"
	"github.com/golang/glog"
)

type Deduct struct {
	Level           int
	TaxedIncomeTop  float64
	DeductRate      float64
	QuickCalcDeduct float64
}

type Request struct {
	MonthlyIncome     float64
	TaxBase           float64
	Insurrance        float64
	AdditionalTaxFree float64
	Months            int
}

type Result struct {
	// 0. 累积应扣税
	// 1. 逐月的扣税
	// 2. 逐月的到手数
	// 3. 合计的到手数
	AccumulatedDeducts []float64
	Deducts            []float64
	Obtains            []float64
	TotalObtains       float64
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

func Calc(req Request) (Result, error) {
	//fmt.Println(req)
	//TODO
	res := Result{
		AccumulatedDeducts: make([]float64, req.Months),
		Deducts:            make([]float64, req.Months),
		Obtains:            make([]float64, req.Months),
		TotalObtains:       0,
	}

	for m := 0; m < req.Months; m++ {
		taxedSalary := (req.MonthlyIncome - req.TaxBase - req.Insurrance - req.AdditionalTaxFree) * (float64(m) + 1)
		glog.V(2).Infoln("taxedSalary", taxedSalary)
		d, err := calcDeductLevel(taxedSalary)
		if nil != err {
			return res, err
		}
		glog.V(2).Infoln("deductLevel", d)
		accumulatedShouldDeduct := taxedSalary*d.DeductRate/100 - d.QuickCalcDeduct
		res.AccumulatedDeducts[m] = accumulatedShouldDeduct
		glog.V(2).Infoln("accumulatedShouldDeduct", accumulatedShouldDeduct)
		currentShouldDeduct := accumulatedShouldDeduct - pastDeducted(res, m)
		glog.V(2).Infoln("currentShouldDeduct", currentShouldDeduct)
		res.Deducts[m] = currentShouldDeduct
		res.Obtains[m] = req.MonthlyIncome - req.Insurrance - currentShouldDeduct

		res.TotalObtains += res.Obtains[m]

		glog.V(1).Infof("%+v\n", res)
	}

	return res, nil
}

func pastDeducted(res Result, month int) float64 {
	if month < 1 {
		return 0
	}
	return res.AccumulatedDeducts[month-1]
}

func calcDeductLevel(taxedSalary float64) (Deduct, error) {
	d := Deduct{}
	for _, d = range deducts {
		if taxedSalary < d.TaxedIncomeTop {
			return d, nil
		}
	}
	return d, errors.New("Taxed salary exceeds max")
}
