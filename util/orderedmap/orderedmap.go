// Package orderedmap implements an ordered map.
//
// Example
//
//	m := orderedmap.NewOrderedMap()
//
//	// Set key with value.
//	m.Set("key", 1)
//
//	// Get value by key.
//	val, ok := m.Get("key")
//
//	// Get keys in insertion order.
//	for _, key := range m.Keys() {
//		val, ok := m.Get(key)
//	}
//
package orderedmap

// OrderedMap is an ordered map with strings as keys.
type OrderedMap struct {
	keys []string
	m    map[string]interface{}
}

// NewOrderedMap returns a new OrderedMap.
func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		keys: make([]string, 0),
		m:    make(map[string]interface{}),
	}
}

// Keys returns the keys in insertion order.
func (m *OrderedMap) Keys() []string { return m.keys }

// Values returns the values in insertion order.
func (m *OrderedMap) Values() []interface{} {
	values := make([]interface{}, 0)
	for _, key := range m.keys {
		values = append(values, m.m[key])
	}
	return values
}

// Set a key into this map.
func (m *OrderedMap) Set(key string, val interface{}) {
	_, ok := m.m[key]
	if !ok {
		m.keys = append(m.keys, key)
	}
	m.m[key] = val
}

// Get a value by key.
func (m *OrderedMap) Get(key string) (val interface{}, ok bool) {
	val, ok = m.m[key]
	return
}

// Pop a value by key.
func (m *OrderedMap) Pop(key string) (val interface{}, ok bool) {
	val, ok = m.m[key]
	delete(m.m, key)
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
	return
}

// Len returns the length of this map.
func (m *OrderedMap) Len() int { return len(m.m) }

// OrderedIntMap is an ordered map with integers as keys.
type OrderedIntMap struct {
	keys []int
	m    map[int]interface{}
}

// NewOrderedIntMap returns a new OrderedIntMap.
func NewOrderedIntMap() *OrderedIntMap {
	return &OrderedIntMap{
		keys: make([]int, 0),
		m:    make(map[int]interface{}),
	}
}

// Keys returns the keys in insertion order.
func (m *OrderedIntMap) Keys() []int { return m.keys }

// Values returns the values in insertion order.
func (m *OrderedIntMap) Values() []interface{} {
	values := make([]interface{}, 0)
	for _, key := range m.keys {
		values = append(values, m.m[key])
	}
	return values
}

// Set a key into this map.
func (m *OrderedIntMap) Set(key int, val interface{}) {
	_, ok := m.m[key]
	if !ok {
		m.keys = append(m.keys, key)
	}
	m.m[key] = val
}

// Get a value by key.
func (m *OrderedIntMap) Get(key int) (val interface{}, ok bool) {
	val, ok = m.m[key]
	return
}

// Pop a value by key.
func (m *OrderedIntMap) Pop(key int) (val interface{}, ok bool) {
	val, ok = m.m[key]
	delete(m.m, key)
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
	return
}

// Len returns the length of this map.
func (m *OrderedIntMap) Len() int { return len(m.m) }

// OrderedUintMap is an ordered map with unsigned integers as keys.
type OrderedUintMap struct {
	keys []uint
	m    map[uint]interface{}
}

// NewOrderedUintMap returns a new OrderedUintMap.
func NewOrderedUintMap() *OrderedUintMap {
	return &OrderedUintMap{
		keys: make([]uint, 0),
		m:    make(map[uint]interface{}),
	}
}

// Keys returns the keys in insertion order.
func (m *OrderedUintMap) Keys() []uint { return m.keys }

// Values returns the values in insertion order.
func (m *OrderedUintMap) Values() []interface{} {
	values := make([]interface{}, 0)
	for _, key := range m.keys {
		values = append(values, m.m[key])
	}
	return values
}

// Set a key into this map.
func (m *OrderedUintMap) Set(key uint, val interface{}) {
	_, ok := m.m[key]
	if !ok {
		m.keys = append(m.keys, key)
	}
	m.m[key] = val
}

// Get a value by key.
func (m *OrderedUintMap) Get(key uint) (val interface{}, ok bool) {
	val, ok = m.m[key]
	return
}

// Pop a value by key.
func (m *OrderedUintMap) Pop(key uint) (val interface{}, ok bool) {
	val, ok = m.m[key]
	delete(m.m, key)
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
	return
}

// Len returns the length of this map.
func (m *OrderedUintMap) Len() int { return len(m.m) }
