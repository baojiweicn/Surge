package device

import (
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func Test_Config(t *testing.T) {
	f, err := os.OpenFile("../../addones/xiaomi/vaccum/vaccum.yml", os.O_RDONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()
	readByte, err := ioutil.ReadAll(f)
	c := &Template{}
	yaml.Unmarshal(readByte, c)
	t.Logf("unmarshal to %+v", c)
	t.Errorf("%+v", c)
}
