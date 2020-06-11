package orderedmap_test

import (
	"testing"

	"github.com/baojiweicn/Surge/util/orderedmap"
	util "github.com/baojiweicn/Surge/util/testutil"
)

func TestOrderedMap(t *testing.T) {
	m := orderedmap.NewOrderedMap()

	m.Set("key1", "val1")
	m.Set("key2", "val2")
	m.Set("key3", "val3")

	val1, ok := m.Get("key1")
	util.MustT(t, ok && val1.(string) == "val1")
	val2, ok := m.Get("key2")
	util.MustT(t, ok && val2.(string) == "val2")
	val3, ok := m.Get("key3")
	util.MustT(t, ok && val3.(string) == "val3")

	keys := m.Keys()
	util.MustT(t, keys[0] == "key1")
	util.MustT(t, keys[1] == "key2")
	util.MustT(t, keys[2] == "key3")

	values := m.Values()

	util.MustT(t, values[0].(string) == "val1")
	util.MustT(t, values[1].(string) == "val2")
	util.MustT(t, values[2].(string) == "val3")

	util.MustT(t, m.Len() == 3)

	val1, ok = m.Pop("key1")
	util.MustT(t, ok && val1.(string) == "val1")
	util.MustT(t, m.Len() == 2)

	val1, ok = m.Get("key1")
	util.MustT(t, !ok)
}

func TestOrderedIntMap(t *testing.T) {
	m := orderedmap.NewOrderedIntMap()

	m.Set(1, "val1")
	m.Set(2, "val2")
	m.Set(3, "val3")

	val1, ok := m.Get(1)
	util.MustT(t, ok && val1.(string) == "val1")
	val2, ok := m.Get(2)
	util.MustT(t, ok && val2.(string) == "val2")
	val3, ok := m.Get(3)
	util.MustT(t, ok && val3.(string) == "val3")

	keys := m.Keys()
	util.MustT(t, keys[0] == 1)
	util.MustT(t, keys[1] == 2)
	util.MustT(t, keys[2] == 3)

	values := m.Values()

	util.MustT(t, values[0].(string) == "val1")
	util.MustT(t, values[1].(string) == "val2")
	util.MustT(t, values[2].(string) == "val3")

	util.MustT(t, m.Len() == 3)

	val1, ok = m.Pop(1)
	util.MustT(t, ok && val1.(string) == "val1")
	util.MustT(t, m.Len() == 2)

	val1, ok = m.Get(1)
	util.MustT(t, !ok)
}

func TestOrderedUintMap(t *testing.T) {
	m := orderedmap.NewOrderedUintMap()

	m.Set(1, "val1")
	m.Set(2, "val2")
	m.Set(3, "val3")

	val1, ok := m.Get(1)
	util.MustT(t, ok && val1.(string) == "val1")
	val2, ok := m.Get(2)
	util.MustT(t, ok && val2.(string) == "val2")
	val3, ok := m.Get(3)
	util.MustT(t, ok && val3.(string) == "val3")

	keys := m.Keys()
	util.MustT(t, keys[0] == 1)
	util.MustT(t, keys[1] == 2)
	util.MustT(t, keys[2] == 3)

	values := m.Values()

	util.MustT(t, values[0].(string) == "val1")
	util.MustT(t, values[1].(string) == "val2")
	util.MustT(t, values[2].(string) == "val3")

	util.MustT(t, m.Len() == 3)

	val1, ok = m.Pop(1)
	util.MustT(t, ok && val1.(string) == "val1")
	util.MustT(t, m.Len() == 2)

	val1, ok = m.Get(1)
	util.MustT(t, !ok)
}
