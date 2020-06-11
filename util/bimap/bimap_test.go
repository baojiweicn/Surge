package bimap_test

import (
	"testing"

	"github.com/baojiweicn/Surge/util/bimap"
	util "github.com/baojiweicn/Surge/util/testutil"
)

func TestBiMap(t *testing.T) {
	bi := bimap.NewBiMap()
	bi.Set(1, "Hello")
	bi.Set(2, "World")

	val1, ok := bi.Get(1)
	val2, ok := bi.Get(2)
	util.MustT(t, ok && val1.(string) == "Hello")
	util.MustT(t, ok && val2.(string) == "World")

	key1, ok := bi.GetKey("Hello")
	key2, ok := bi.GetKey("World")
	util.MustT(t, ok && key1.(int) == 1)
	util.MustT(t, ok && key2.(int) == 2)

	length := bi.Len()
	util.MustT(t, length == 2)

	pop1, ok := bi.Pop(1)
	util.MustT(t, pop1.(string) == "Hello")

	keys := bi.Keys()
	util.MustT(t, ok && keys[0].(int) == 2)
	values := bi.Values()
	util.MustT(t, values[0].(string) == "World")
}
