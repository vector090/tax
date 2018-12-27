package lib

import (
	"flag"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	TAX_BASE = 5000.0
)

func TestAlgo(t *testing.T) {
	flag.Set("logtostderr", "true")
	flag.Set("v", "3")

	req := Request{
		MonthlyIncome:     15000,
		TaxBase:           TAX_BASE,
		Insurrance:        3200,
		AdditionalTaxFree: 3500,
		Months:            1,
	}

	res, err := Calc(req)
	Convey("Calc example", t, func() {
		So(err, ShouldBeNil)
		So(res.Deducts, ShouldHaveLength, 1)
		So(res.Deducts[0], ShouldEqual, 99)
		//So(res.Deducts, ShouldAlmostEqual, 99)
	})

	Convey("Calc test", t, func() {
		req = Request{
			MonthlyIncome:     25000,
			TaxBase:           TAX_BASE,
			Insurrance:        532.14,
			AdditionalTaxFree: 3000,
			Months:            1,
		}

		runPassedTest := false
		if runPassedTest {
			res, err = Calc(req)
			So(err, ShouldBeNil)
			So(res.Deducts, ShouldHaveLength, 1)
			So(res.Deducts[0], ShouldEqual, 494.0358)

			req.Months = 2
			res, err = Calc(req)
			So(err, ShouldBeNil)
			So(res.Deducts, ShouldHaveLength, 2)
			So(res.Deducts[0], ShouldEqual, 494.0358)
			So(res.Deducts[1], ShouldEqual, 494.0358)
		}

		req.Months = 3
		res, err = Calc(req)
		So(err, ShouldBeNil)
		So(res.Deducts, ShouldHaveLength, 3)
		So(res.Deducts[0], ShouldEqual, 494.0358)
		So(res.Deducts[1], ShouldEqual, 494.0358)
		So(res.Deducts[2], ShouldAlmostEqual, 1432.2864)
	})
}

func TestCalcDeductLevel(t *testing.T) {
	Convey("calcDeductLevel", t, func() {
		d, err := calcDeductLevel(35000)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[0])

		d, err = calcDeductLevel(0)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[0])

		d, err = calcDeductLevel(-1)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[0])

		d, err = calcDeductLevel(36000 - 1)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[0])

		d, err = calcDeductLevel(36000)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[1])

		d, err = calcDeductLevel(144000 - 1)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[1])

		d, err = calcDeductLevel(144000)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[2])

		d, err = calcDeductLevel(300000 - 1)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[2])

		d, err = calcDeductLevel(300000)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[3])

		d, err = calcDeductLevel(420000 - 1)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[3])

		d, err = calcDeductLevel(420000)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[4])

		d, err = calcDeductLevel(660000 - 1)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[4])

		d, err = calcDeductLevel(660000)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[5])

		d, err = calcDeductLevel(960000 - 1)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[5])

		d, err = calcDeductLevel(960000)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[6])

		d, err = calcDeductLevel(1000000)
		So(err, ShouldBeNil)
		So(d, ShouldResemble, deducts[6])

		d, err = calcDeductLevel(9999990000 + 1)
		So(err, ShouldNotBeNil)
		//So(d, ShouldResemble, deducts[6])
	})
}
