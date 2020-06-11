package hash_test

import (
	"testing"

	"github.com/baojiweicn/Surge/util/hash"
)

func TestCrc16(t *testing.T) {
	res := hash.Crc16([]byte("123456789"))

	if res != 0x4B37 {
		t.Errorf("Expected 0x4B37 actual %d", res)
	}

	if 0x32E4 != hash.Crc16([]byte("123456")) {
		t.Errorf("Expected 0x32E4")
	}
}
