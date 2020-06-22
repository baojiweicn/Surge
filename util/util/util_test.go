package util_test

import (
	"testing"

	tutil "github.com/baojiweicn/Surge/util/testutil"
	util "github.com/baojiweicn/Surge/util/util"
)

type Duck interface {
	GuaGuaGua()
}

type YellowDuck struct{}

func (d *YellowDuck) GuaGuaGua() {}

func TestIsNilInterface(t *testing.T) {
	var d Duck
	var d1 *YellowDuck
	d = d1
	tutil.MustT(t, d != nil)
	tutil.MustT(t, util.IsNilInterface(d))
}

func TestGenerateRandString(t *testing.T) {
	tutil.MustT(t, util.GenerateRandString(10) != util.GenerateRandString(10))
	tutil.MustT(t, util.GenerateRandString(3) != util.GenerateRandString(4))
}
