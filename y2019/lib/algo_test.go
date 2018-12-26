package lib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	TAX_BASE = 5000.0
)

func TestAlgo(t *testing.T) {
	req := Request{
		MonthlyIncome: 15000,
		TaxBase:       TAX_BASE,
		Insurrance:    3200,
		Months:        1,
	}
	Convey("Calc 1", t, func() {
		res, err := Calc(req)
		So(err, ShouldBeNil)
		So(res.Deducts, ShouldHaveLength, 1)
		So(res.Deducts[0], ShouldEqual, 99)
		//So(res.Deducts, ShouldAlmostEqual, 99)
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
