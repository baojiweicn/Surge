package lock_test

import (
	"testing"
	"time"

	"github.com/baojiweicn/Surge/util/lock"
	util "github.com/baojiweicn/Surge/util/testutil"
)

func TestMutex(t *testing.T) {
	defaultLogErrorFunc := lock.DefaultLogFunc
	m := lock.NewMutex("aMutex", 100*time.Millisecond)
	isCalled := false
	lock.LogFunc(func(m lock.MutexWrapper, op string) {
		isCalled = true
		defaultLogErrorFunc(m, op)
	})
	go func() {
		m.Lock()
		defer m.Unlock()
		time.Sleep(120 * time.Millisecond)
	}()
	m.Lock()
	defer m.Unlock()
	time.Sleep(120 * time.Millisecond)
	util.MustT(t, isCalled)
}

func TestRWMutex(t *testing.T) {
	defaultLogErrorFunc := lock.DefaultLogFunc
	m := lock.NewRWMutex("aRWMutex", 100*time.Millisecond)
	isCalled := false
	lock.LogFunc(func(m lock.MutexWrapper, op string) {
		isCalled = true
		defaultLogErrorFunc(m, op)
	})
	go func() {
		m.Lock()
		defer m.Unlock()
		time.Sleep(120 * time.Millisecond)
	}()
	m.Lock()
	defer m.Unlock()
	time.Sleep(120 * time.Millisecond)
	util.MustT(t, isCalled)
}

func TestRWMutexBlockRlock(t *testing.T) {
	defaultLogErrorFunc := lock.DefaultLogFunc
	m := lock.NewRWMutex("aRWMutex", 100*time.Millisecond)
	isCalled := false
	lock.LogFunc(func(m lock.MutexWrapper, op string) {
		isCalled = true
		defaultLogErrorFunc(m, op)
	})
	go func() {
		m.RLock()
		defer m.RUnlock()
		time.Sleep(120 * time.Millisecond)
	}()
	m.Lock()
	defer m.Unlock()
	time.Sleep(120 * time.Millisecond)
	util.MustT(t, isCalled)
}
