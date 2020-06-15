package source

import (
	"fmt"
	"testing"
)

func Test_PythonManager(t *testing.T) {
	manager := GetDefaultPythonManager()
	if manager != nil {
		fmt.Println(manager.Path())
		fmt.Print(manager.GetAll())
	} else {
		fmt.Print("nil")
	}
}
