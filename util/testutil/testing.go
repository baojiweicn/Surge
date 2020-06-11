// Package testutil provides testing utils.
package testutil

import (
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/baojiweicn/Surge/util/log"
)

// MustT asserts the given value is True for testing.
func MustT(t *testing.T, v bool) {
	Must(t, nil, v)
}

// MustT asserts the given value is True for benchmark.
func MustB(b *testing.B, v bool) {
	Must(nil, b, v)
}

// Must asserts the given value is True for testing or benchmark.
func Must(t *testing.T, b *testing.B, v bool) {
	if !v {
		_, fileName, line, _ := runtime.Caller(2)
		if t != nil {
			t.Error(log.Colored("red", "\n ======================================= TEST FAIL [START] ========================================"))
			t.Error(log.Colored("red", fmt.Sprintf("\n unexcepted: %s:%d", fileName, line)))
			t.Errorf("\n %s", debug.Stack())
			t.Error(log.Colored("red", "\n ======================================= TEST FAIL [END] =========================================="))
			t.FailNow()
		}
		if b != nil {
			b.Error(log.Colored("red", "\n ======================================= TEST FAIL [START] ========================================"))
			b.Error(log.Colored("red", fmt.Sprintf("\n unexcepted: %s:%d", fileName, line)))
			b.Errorf("\n %s", debug.Stack())
			b.Error(log.Colored("red", "\n ======================================= TEST FAIL [END] =========================================="))
			b.FailNow()
		}
	}
}

// getTestFunc returns the function named `name` from value `xv`.
func getTestFunc(t *testing.T, xv reflect.Value, name string) func(*testing.T) {
	if m := xv.MethodByName(name); m.IsValid() {
		if f, ok := m.Interface().(func(*testing.T)); ok {
			return f
		}
		// Method exists but has the wrong type signature.
		t.Fatalf("grpctest: function %v has unexpected signature (%T)", name, m.Interface())
	}
	return func(*testing.T) {}
}

// RunSubTests runs all subtests bind on a test group s with case-level Setup
// and Teardown support.
//
// Steps to use subtests.
//
//	1. Declares an empty struct
//		type s struct{}
//	2. Implements Setup() and Teardown() on s.
//		func (s) Setup(t *testing.T) {}
//		func (s) Teardown(t *testing.T) {}
//	3. Writes your tests..
//		func (s) TestMyFeature(t *testing.T) {}
//	4. Write a top-level test case to trigger all subtests.
//		func TestAllSubTests(t *testing.T) {
//			testutil.RunSubTests(t, s{})
//		}
//
// Orginal from
// https://github.com/grpc/grpc-go/blob/master/internal/grpctest/grpctest.go
func RunSubTests(t *testing.T, x interface{}) {
	xt := reflect.TypeOf(x)
	xv := reflect.ValueOf(x)

	setup := getTestFunc(t, xv, "Setup")
	teardown := getTestFunc(t, xv, "Teardown")

	for i := 0; i < xt.NumMethod(); i++ {
		methodName := xt.Method(i).Name
		if !strings.HasPrefix(methodName, "Test") {
			continue
		}
		tfunc := getTestFunc(t, xv, methodName)
		t.Run(strings.TrimPrefix(methodName, "Test"), func(t *testing.T) {
			setup(t)
			// defer teardown to guarantee it is run even if tfunc uses t.Fatal()
			defer teardown(t)
			tfunc(t)
		})
	}
}
