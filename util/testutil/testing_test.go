package testutil_test

import (
	"testing"

	"github.com/baojiweicn/Surge/util/testutil"
)

type s struct{}

// Global states.
var globalState = 1

func (s) Setup(t *testing.T) {
	globalState = 2
}

func (s) Teardown(t *testing.T) {
	// Rollback
	globalState = 1
}

func (s) TestA(t *testing.T) {
	testutil.MustT(t, globalState == 2)
}

func (s) TestB(t *testing.T) {
	testutil.MustT(t, globalState == 2)
}

func Test(t *testing.T) {
	testutil.RunSubTests(t, s{})
}
