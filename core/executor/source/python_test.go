package source

import (
	"testing"
)

func Test_PythonManager(t *testing.T) {
	manager := GetDefaultPythonManager()
	if manager != nil {
		t.Logf(manager.Path())
		if all, err := manager.GetAll(); err == nil {
			t.Log(all[0])
		}
	} else {
		t.Errorf("nil manager")
	}
}
