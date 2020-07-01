package source

import (
	"testing"
)

func Test_PythonManager(t *testing.T) {
	manager := GetDefaultPythonManager()
	if manager != nil {
		t.Logf(manager.Path())
	} else {
		t.Errorf("nil manager")
	}
}

func Test_PythonManageGetAll(t *testing.T) {
	manager := GetDefaultPythonManager()
	if manager != nil {
		t.Logf(manager.Path())
		if all, err := manager.GetAll(); err != nil {
			t.Error(err)
		} else {
			for _, p := range all {
				t.Log(p)
			}
		}
	} else {
		t.Errorf("nil manager")
	}

}

func Test_PythonManageInstall(t *testing.T) {
	manager := GetDefaultPythonManager()
	pack := &Package{
		Name: "requests",
	}
	defer manager.Uninstall(pack)
	if manager != nil {
		t.Logf(manager.Path())
		if err := manager.Install(pack); err != nil {
			t.Error(err)
		}
	}
}
